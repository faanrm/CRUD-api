package routes

import (
	"github.com/faanrm/CRUD-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(routes *gin.RouterGroup) {
	userController := controllers.UserControllers()
	routes.POST("/users", userController.CreateUser)
	routes.GET("/users", userController.GetUser)
	routes.PUT("/users/:userId", userController.CreateUser)
	routes.GET("/users/:userId", userController.GetUser)
	routes.DELETE("/users/:userId", userController.DeleteUser)
}
