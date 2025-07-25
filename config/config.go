package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		panic("failed to get database instance")
	}
	db.Close()
	fmt.Println("Database connection closed")
}
