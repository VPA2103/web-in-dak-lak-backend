package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {

	router.GET("/products", controllers.GetProduct)
	// Tạo user mới
	router.POST("/products/add", controllers.CreateProduct)
	////update user
	//router.PUT("/users/:id/update", controllers.UpdateUser)
	////
	//router.GET("/users/:id", controllers.GetUserIndex)
	router.DELETE("/products/:id", controllers.DeleteProduct)
	//
	router.GET("/products/:id", controllers.GetProductIndex)

}
