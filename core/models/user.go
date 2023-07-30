package models

import "gorm.io/gorm"

type User struct {
	Email         string `gorm:"unique;type:varchar(320);column:email"`
	Password      string `gorm:"column:password;type:varchar(50)"`
	Nickname      string `gorm:"unique;column:nickname;type:varchar(50)"`
	HeadPortrait  string `gorm:"column:head_portrait;type:varchar(200)"`
	CurCapacity   uint   `gorm:"column:cur_capacity"`
	TotalCapacity uint   `gorm:"column:total_capacity"`
	gorm.Model
}

func (u User) TableName() string {
	return "user"
}
