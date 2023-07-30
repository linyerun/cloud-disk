package models

import (
	"gorm.io/gorm"
	"time"
)

type ShareFile struct {
	gorm.Model
	UserId      uint      `gorm:"column:user_id"`
	FileId      uint      `gorm:"column:file_id"`
	ExpiredTime time.Time `gorm:"column:expired_time"`
	ClickNum    uint      `gorm:"column:click_num"`
}

func (ShareFile) TableName() string {
	return "share_file"
}
