package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)

	// Tạo user mới
	router.POST("/users", controllers.CreateUser)

}
