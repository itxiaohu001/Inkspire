package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`           // 用户 ID
	Username  string    `gorm:"uniqueIndex;size:100"` // 用户名，唯一
	Email     string    `gorm:"uniqueIndex;size:100"` // 邮箱，唯一
	Password  string    `gorm:"size:255"`             // 密码（加密存储）
	Role      string    `gorm:"size:50"`              // 角色（admin, user）
	CreatedAt time.Time `gorm:"autoCreateTime"`       // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`       // 更新时间
}
