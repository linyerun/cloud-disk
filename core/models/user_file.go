package models

import "gorm.io/gorm"

type UserFile struct {
	gorm.Model
	UserId   uint   `gorm:"column:user_id"`
	ParentId uint   `gorm:"column:parent_id"`                  // 0表示它处于根目录
	FileId   uint   `gorm:"column:file_id"`                    // 文件夹的file_id为0
	FileType uint8  `gorm:"column:file_type"`                  // 文件夹或者文件
	FileName string `gorm:"type:varchar(255);column:filename"` // 文件或者文件夹用户的自定义取名
}

func (UserFile) TableName() string {
	return "user_file"
}
