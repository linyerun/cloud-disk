package db

import "cloud-disk/core/models"

var tableNameShareFile = models.ShareFile{}.TableName()

func SaveShareFile(shareFile *models.ShareFile) error {
	return MySQLClient.Create(shareFile).Error
}

func GetShareFileById(id uint) (shareFile *models.ShareFile, err error) {
	shareFile = new(models.ShareFile)
	err = MySQLClient.Table(tableNameShareFile).Where("id = ?", id).Find(shareFile).Error
	return
}

func UpdateShareFieldClickNumById(id, clickNum uint) error {
	return MySQLClient.Table(tableNameShareFile).Where("id = ?", id).Update("click_num", clickNum).Error
}
