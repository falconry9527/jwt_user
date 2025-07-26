package main

import (
	"jwt_user/db"
	"jwt_user/routes"
	"log"
)

func main() {
	// 初始化数据库连接
	db.InitDB()
	defer db.CloseDB()

	db.InitRedis()
	defer db.CloseRedis()

	// 初始化路由
	router := routes.InitRouter()

	// 启动服务器
	log.Fatal(router.Run(":8080"))
}
