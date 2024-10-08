package main

import (
	"fmt"
	"net/http"

	"github.com/faanrm/CRUD-api/config"
	"github.com/faanrm/CRUD-api/models"
	"github.com/faanrm/CRUD-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Users{})
	if err != nil {
		fmt.Println("Failed to migrate database:", err)
		return
	}
	fmt.Println("Database migrated succesfully")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	api := r.Group("/api")
	routes.RouteAll(api)
	r.Run(":5200")
}
