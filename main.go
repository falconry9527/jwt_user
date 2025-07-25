package main

import (
	"jwt_user/config"
	"jwt_user/routes"
	"log"
)

func main() {
	// 初始化数据库连接
	config.InitDB()
	defer config.CloseDB()

	// 设置路由
	router := routes.SetupRouter()

	// 启动服务器
	log.Fatal(router.Run(":8080"))
}
