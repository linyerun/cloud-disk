package utils

import (
	"cloud-disk/core/define"
	"context"
	"errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"path"
)

// TencentUploadFileByRequest 把Request上的File上传文件到腾讯云
func TencentUploadFileByRequest(r *http.Request) (fileURL string, err error) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	return TencentUploadFile(file, fileHeader.Filename)
}

// TencentUploadFile 上传文件到腾讯云
func TencentUploadFile(file io.Reader, filename string) (fileURL string, err error) {
	client, err := getCosClient()
	if err != nil {
		return
	}
	path := UUID() + path.Ext(filename)
	response, err := client.Object.Put(context.Background(), path, file, nil)
	if err != nil {
		return
	}
	Logger().Println(response.StatusCode, response.Body)
	return define.CosBucket + "/" + path, nil
}

func TencentDownloadFile(fileURL string) (resp *cos.Response, err error) {
	client, err := getCosClient()
	if err != nil {
		return nil, err
	}
	if len(fileURL) < len(define.CosBucket+"/") {
		return nil, errors.New("fileURL错误")
	}
	resp, err = client.Object.Get(context.Background(), fileURL[len(define.CosBucket+"/"):], nil)
	return
}

func getCosClient() (client *cos.Client, err error) {
	cosURL, err := url.Parse(define.CosBucket)
	if err != nil {
		return
	}
	baseURL := &cos.BaseURL{
		BucketURL: cosURL,
	}
	client = cos.NewClient(baseURL, &http.Client{Transport: &cos.AuthorizationTransport{
		SecretID:  define.TencentSecretID,
		SecretKey: define.TencentSecretKey,
	}})
	return
}
