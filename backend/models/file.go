package models

import "time"

type File struct {
	ID         uint      `gorm:"primaryKey"`        // 文件 ID
	FileName   string    `gorm:"size:255;not null"` // 文件名称
	FilePath   string    `gorm:"size:255;not null"` // 文件路径
	FileType   string    `gorm:"size:50"`           // 文件类型（如 image/png, application/pdf）
	UploadedBy uint      `gorm:"index"`             // 上传者 ID，关联用户表
	CreatedAt  time.Time `gorm:"autoCreateTime"`    // 上传时间
}
