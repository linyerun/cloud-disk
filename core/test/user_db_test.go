package test

import (
	"cloud-disk/core/db"
	"fmt"
	"testing"
)

func init() {
	db.InitMySQLForTest("root", "123456", "localhost", "cloud-disk", 3306)
}

func TestGetUserByEmail(t *testing.T) {
	user, err := db.GetUserByEmail("2338244917@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", user)
}

func TestGetUserCapacityById(t *testing.T) {
	user, err := db.GetUserCapacityById(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user.TotalCapacity, user.CurCapacity)
}
