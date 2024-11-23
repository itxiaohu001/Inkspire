package models

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey"`          // 博客文章 ID
	Title       string    `gorm:"size:255;not null"`   // 标题
	Content     string    `gorm:"type:text;not null"`  // 正文内容（支持 Markdown）
	AuthorID    uint      `gorm:"index"`               // 作者 ID，关联用户表
	CategoryID  uint      `gorm:"index"`               // 分类 ID，关联分类表
	Tags        []*Tag    `gorm:"many2many:post_tags"` // 标签（多对多）
	ViewCount   uint      `gorm:"default:0"`           // 浏览量
	CreatedAt   time.Time `gorm:"autoCreateTime"`      // 创建时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`      // 更新时间
	PublishedAt time.Time `gorm:"index"`               // 发布时间
	IsDraft     bool      `gorm:"default:false"`       // 是否为草稿
}
