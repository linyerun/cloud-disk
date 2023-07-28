package db

import (
	"cloud-disk/core/models"
	"cloud-disk/core/utils"
)

func HasRegistered(email string) bool {
	var cnt int64
	err := MySQLClient.Table("user").Where("email = ?", email).Count(&cnt).Error
	if err != nil {
		utils.Logger().Error(err)
		return false
	}
	return cnt > 0
}

func SaveUser(user *models.User) error {
	return MySQLClient.Create(user).Error
}

func HasTheNickname(nickname string) bool {
	var cnt int64
	err := MySQLClient.Table("user").Where("nickname = ?", nickname).Count(&cnt).Error
	if err != nil {
		utils.Logger().Error(err)
		return false
	}
	return cnt > 0
}

func HasTheUser(email string, password string) bool {
	var cnt int64
	err := MySQLClient.Table("user").Where("email = ? and password = ?", email, password).Count(&cnt).Error
	if err != nil {
		return false
	}
	return cnt > 0
}
