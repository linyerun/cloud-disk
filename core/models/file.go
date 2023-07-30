package models

import "gorm.io/gorm"

// File 真正存储实打实文件的地方
type File struct {
	gorm.Model
	Hash string `gorm:"type:varchar(32);column:hash"`
	Size uint   `gorm:"column:size"`
	Path string `gorm:"type:varchar(255);column:path"`
}

func (File) TableName() string {
	return "file"
}
