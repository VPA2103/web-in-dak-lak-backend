package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func UploadRoutes(route *gin.Engine) {
	route.POST("/upload", controllers.UploadHandler)
	route.GET("/images", controllers.GetProductImage)
}
