package routes

import (
	"github.com/gin-gonic/gin"
)

func RouteAll(routes *gin.RouterGroup) {
	UserRoute(routes)
}
