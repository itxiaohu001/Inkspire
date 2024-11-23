package models

import "time"

type UserStatistic struct {
	ID        uint      `gorm:"primaryKey"`     // 记录 ID
	UserID    uint      `gorm:"index"`          // 用户 ID
	PostID    uint      `gorm:"index"`          // 文章 ID
	Action    string    `gorm:"size:50"`        // 行为类型（view, like, share）
	CreatedAt time.Time `gorm:"autoCreateTime"` // 行为发生时间
}
