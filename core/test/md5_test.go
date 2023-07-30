package test

import (
	"cloud-disk/core/utils"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	pwd := utils.Md5("123456")
	fmt.Println(pwd)
}
