package models

import (
	"gorm.io/gorm"
)

type ShareFile struct {
	gorm.Model
	UserId      uint  `gorm:"column:user_id"`
	FileId      uint  `gorm:"column:file_id"`
	ExpiredTime int64 `gorm:"column:expired_time"` // -1: 没有过期时间, 非零则表示过期时间的时间戳
	ClickNum    uint  `gorm:"column:click_num"`
}

func (ShareFile) TableName() string {
	return "share_file"
}
