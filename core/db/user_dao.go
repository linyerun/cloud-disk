package db

import (
	"cloud-disk/core/models"
	"cloud-disk/core/utils"
	"errors"
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

func HasTheUser(email string, password string) (uint, bool) {
	var id uint
	err := MySQLClient.Table("user").Select("id").Where("email = ? and password = ?", email, password).Scan(&id).Error
	if err != nil {
		utils.Logger().Error(err)
		return 0, false
	}
	return id, id > 0
}

func GetUserByEmail(email string) (user *models.User, err error) {
	user = &models.User{}
	var cnt int64
	err = MySQLClient.Table("user").Select("email, nickname, head_portrait").Where("email = ?", email).First(user).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	if cnt < 1 {
		return nil, errors.New("用户不存在")
	}
	return user, nil
}

func GetUserCapacityById(uid int) (user *models.User, err error) {
	user = &models.User{}
	err = MySQLClient.Table("user").Select("cur_capacity, total_capacity").Where("id = ?", uid).First(user).Error
	return
}
