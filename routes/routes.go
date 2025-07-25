package routes

import (
	"github.com/gin-gonic/gin"
	"jwt_user/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userController := controllers.NewUserController()

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", userController.GetAllUsers)
			users.POST("/", userController.CreateUser)
			users.GET("/:id", userController.GetUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}

	return r
}
