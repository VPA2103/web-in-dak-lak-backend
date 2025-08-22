package main

import (
	"backend/config"
	_ "backend/config"
	"backend/models"

	_ "backend/models"

	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	// Tạo router mặc định
	r := gin.Default()

	// Route GET cơ bản
	routes.UserRoute(r)

	// Chạy server trên port 8080
	r.Run(":8080")
}
