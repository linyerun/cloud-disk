package process

import (
	"cloud-disk/core/define"
	"math/rand"
	"time"
)

var (
	chs []chan func()
)

func AddTask(task func()) {
	rand.Seed(time.Now().Unix())
	chs[rand.Uint64()%define.ProcessingCenterPoolSize] <- task
}

func Init() {
	chs = make([]chan func(), define.ProcessingCenterPoolSize)
	for i := uint(0); i < define.ProcessingCenterPoolSize; i++ {
		chs[i] = make(chan func(), define.ProcessingCenterChanLen)
		go processHandler(chs[i])
	}
}

func processHandler(ch chan func()) {
	for {
		handle := <-ch
		handle()
	}
}
