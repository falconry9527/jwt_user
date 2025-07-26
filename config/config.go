package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwt_user/models"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	dsn := "root:Conry@1238@tcp(127.0.0.1:3306)/jwt?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.User{})
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
