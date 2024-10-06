package main

import (
	"net/http"

	"github.com/faanrm/CRUD-api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	_ = db
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	r.Run(":5200")
}
