package models

import "time"

type Tag struct {
	ID        uint      `gorm:"primaryKey"`      // 标签 ID
	Name      string    `gorm:"size:100;unique"` // 标签名称，唯一
	CreatedAt time.Time `gorm:"autoCreateTime"`  // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`  // 更新时间
}
