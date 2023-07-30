package db

import (
	"cloud-disk/core/models"
	"errors"
)

func GetFileByHash(hash string) (fid uint, err error) {
	err = MySQLClient.Table("file").Select("id").Where("hash = ?", hash).Scan(&fid).Error
	if err == nil && fid == 0 {
		err = errors.New("不存在该文件")
	}
	return
}

func SaveFile(file *models.File) error {
	return MySQLClient.Create(file).Error
}
