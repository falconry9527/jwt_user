package routes

import (
	"github.com/gin-gonic/gin"
	"jwt_user/controllers"
	"jwt_user/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
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
