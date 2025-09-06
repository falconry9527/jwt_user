package db

import (
	"fmt"
	"jwt_user/config"
	"time"

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
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// Logger： 设置日志打印sql
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed to use database")
	}
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns)                                // 空闲连接数   默认最大2个空闲连接数  使用默认值即可
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns)                                // 最大连接数   默认0是无限制的  使用默认值即可
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxLifeTime) * time.Second) // 最大等待时间 0 是无穷大
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
