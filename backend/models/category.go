package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`      // 分类 ID
	Name      string    `gorm:"size:100;unique"` // 分类名称，唯一
	CreatedAt time.Time `gorm:"autoCreateTime"`  // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`  // 更新时间
}
