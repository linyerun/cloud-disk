package db

import (
	"cloud-disk/core/models"
	"errors"
)

var tableNameFile = models.File{}.TableName()

func GetFileByHash(hash string) (fid uint, err error) {
	err = MySQLClient.Table(tableNameFile).Select("id").Where("hash = ?", hash).Scan(&fid).Error
	if err == nil && fid == 0 {
		err = errors.New("不存在该文件")
	}
	return
}

func SaveFile(file *models.File) error {
	return MySQLClient.Create(file).Error
}

func GetFileById(id uint) (file *models.File, err error) {
	file = new(models.File)
	err = MySQLClient.Table(tableNameFile).Where("id = ?", id).Find(&file).Error
	return
}
