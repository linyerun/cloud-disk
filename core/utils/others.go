package utils

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
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
