package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {

	router.GET("/categories", controllers.GetCategories)
	// Tạo user mới
	router.POST("/category/add", controllers.CreateCategory)
	////update user
	//router.PUT("/users/:id/update", controllers.UpdateUser)
	////
	//router.GET("/users/:id", controllers.GetUserIndex)
	router.DELETE("/category/:id", controllers.DeleteaCategory)

}
