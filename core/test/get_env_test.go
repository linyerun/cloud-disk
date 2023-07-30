package test

import (
	"cloud-disk/core/define"
	"fmt"
	"testing"
)

func TestGetEnv(t *testing.T) {
	fmt.Println(define.MailPassword)
	fmt.Println(define.CosBucket)
	fmt.Println(define.TencentSecretID)
	fmt.Println(define.TencentSecretKey)
}
