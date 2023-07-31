package models

import (
	"gorm.io/gorm"
)

type ShareFile struct {
	gorm.Model
	UserId      uint  `gorm:"column:user_id"`
	FileId      uint  `gorm:"column:file_id"`
	ExpiredTime int64 `gorm:"column:expired_time"`
	ClickNum    uint  `gorm:"column:click_num"`
}

func (ShareFile) TableName() string {
	return "share_file"
}
