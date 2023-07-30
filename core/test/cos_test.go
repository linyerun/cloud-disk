package test

import (
	"cloud-disk/core/utils"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestTencentFileUpload(t *testing.T) {
	file, err := os.Open("cos_test.go")
	if err != nil {
		t.Fatal(err)
	}
	url, err := utils.TencentUploadFile(file, "cos_test.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("url:", url)
}

func TestTencentFileDownload(t *testing.T) {
	resp, err := utils.TencentDownloadFile("https://cloud-disk-1308170155.cos.ap-guangzhou.myqcloud.com/62d35b6c-babb-45f4-9e8d-b5a811caf571.go")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(utils.Hash(bytes))
}
