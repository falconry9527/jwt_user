package db

import (
	"fmt"
	"jwt_user/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	mysqlConf := config.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.UserName,
		mysqlConf.Password,
		mysqlConf.Address,
		mysqlConf.Port,
		mysqlConf.DbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	// Logger： 设置日志打印sql
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
