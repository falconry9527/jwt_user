package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jwt_user/controllers"
	"jwt_user/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 初始化日志记录器
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	// 中间件鉴权和打印日志
	router.Use(middleware.JWTAuthMiddleware(), middleware.RequestLoggerMiddleware(logger))
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/getAll", controllers.GetAllUsers)
			users.POST("/create", controllers.CreateUser)
			users.GET("/get", controllers.GetUser)
			users.POST("/update", controllers.UpdateUser)
			users.POST("/delete", controllers.DeleteUser)
		}
	}
	return router
}
