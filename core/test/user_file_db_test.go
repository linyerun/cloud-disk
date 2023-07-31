package test

import (
	"cloud-disk/core/db"
	"fmt"
	"testing"
)

func init() {
	db.InitMySQLForTest("root", "123456", "localhost", "cloud-disk", 3306)
}

func TestGetUsers(t *testing.T) {
	files, err := db.GetUserFileListByParentIdUserId(0, 1)
	if err != nil {
		t.Fatal(err)
	}
	for idx, file := range files {
		fmt.Println(idx, file)
	}
}

func TestUpdateUserFile(t *testing.T) {
	err := db.UpdateUserFileName(1, 1, "my_test.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("OK")
}

func TestDeleteUserFileById(t *testing.T) {
	err := db.DeleteUserFileById(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("OK")
}
