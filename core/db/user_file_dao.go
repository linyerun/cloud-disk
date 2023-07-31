package db

import "cloud-disk/core/models"

var tableNameUserFile = models.UserFile{}.TableName()

func HasTheDir(parentId, userId uint) bool {
	if parentId == 0 {
		return true
	} else if parentId < 0 {
		return false
	}
	var cnt int64
	err := MySQLClient.Table(tableNameUserFile).Where("user_id = ? and parent_id = ? and file_type = 0", userId, parentId).Count(&cnt).Error
	if err != nil {
		return false
	}
	return cnt > 0
}

func SaveUserFile(userFile *models.UserFile) error {
	return MySQLClient.Create(userFile).Error
}

func GetUserFileListByParentIdUserId(parentId, userId uint) (userFiles []*models.UserFile, err error) {
	err = MySQLClient.Table(tableNameUserFile).Select("file_id, filename, file_type").Where("parent_id = ? and user_id = ?", parentId, userId).Find(&userFiles).Error
	for _, file := range userFiles {
		file.ParentId = parentId
	}
	return
}

func UpdateUserFileName(userId, userFileId uint, filename string) error {
	return MySQLClient.Table(tableNameUserFile).Where("user_id = ? and id = ?", userId, userFileId).Update("filename", filename).Error
}

func UpdateUserFileParentId(userId, userFileId, parentId uint) error {
	return MySQLClient.Table(tableNameUserFile).Where("user_id = ? and id = ?", userId, userFileId).Update("parent_id", parentId).Error
}

func DeleteUserFileById(userFileId, userId uint) error {
	return MySQLClient.Where("user_id = ? and id = ?", userId, userFileId).Delete(&models.UserFile{}).Error
}
