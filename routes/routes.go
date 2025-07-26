package routes

import (
	"github.com/gin-gonic/gin"
	"jwt_user/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
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

	return r
}
