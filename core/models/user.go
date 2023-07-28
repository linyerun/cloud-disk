package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Email        string         `gorm:"primaryKey;type:varchar(320);column:email"`
	Password     string         `gorm:"column:password;type:varchar(50)"`
	Nickname     string         `gorm:"unique;column:nickname;type:varchar(50)"`
	HeadPortrait string         `gorm:"column:head_portrait;type:varchar(200)"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u User) TableName() string {
	return "user"
}
