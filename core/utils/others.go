package utils

import (
	"crypto/md5"
	"fmt"
)

type H map[string]any

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}