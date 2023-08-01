package utils

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type H map[string]any

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func UUID() string {
	return uuid.NewV4().String()
}

func Hash(bs []byte) string {
	return fmt.Sprintf("%x", md5.Sum(bs))
}

func ToUint(s string) uint {
	num, _ := strconv.Atoi(s)
	return uint(num)
}

func ToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
