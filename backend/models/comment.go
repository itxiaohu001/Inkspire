package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey"`         // 评论 ID
	PostID    uint      `gorm:"index"`              // 博客文章 ID，关联文章表
	UserID    uint      `gorm:"index"`              // 用户 ID，关联用户表
	Content   string    `gorm:"type:text;not null"` // 评论内容
	ParentID  *uint     `gorm:"index"`              // 父评论 ID（用于嵌套评论）
	CreatedAt time.Time `gorm:"autoCreateTime"`     // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`     // 更新时间
}
