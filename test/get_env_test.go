package test

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	fmt.Println(os.Getenv("MailPassword"))
}
