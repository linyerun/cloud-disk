package test

import (
	"cloud-disk/core/db"
	"cloud-disk/core/models"
	"fmt"
	"testing"
)

func init() {
	db.InitMySQLForTest("root", "123456", "localhost", "cloud-disk", 3306)
}

func TestSaveFile(t *testing.T) {
	file := &models.File{
		Hash: "17ef025c98477888a86cd89d89742ca0",
		Size: 311,
		Path: "https://cloud-disk-1308170155.cos.ap-guangzhou.myqcloud.com/62d35b6c-babb-45f4-9e8d-b5a811caf571.go",
	}
	err := db.SaveFile(file)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFile(t *testing.T) {
	fid, err := db.GetFileByHash("17ef025c98477888a86cd89d89742ca0")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("fid:", fid)
}
