package database

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB 初始化数据库连接
func InitDB(dsn string) {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
	})
}

// GetDB 返回全局数据库实例
func GetDB() *gorm.DB {
	return db
}
