package main

import (
	"backend/auth"
	"backend/config"
	_ "backend/config"
	"backend/models"
	_ "backend/models"

	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.
	err := auth.InitKeycloak()
	if err != nil {
		panic(err)
	}

	//2. Tạo router mặc định
	r := gin.Default()
	config.SetupCORS(r)

	//
	config.ConnectDB()
	config.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.News{},
		&models.Contact{},
		&models.Slide{},
		&models.ProductImage{},
		&models.ProductSpec{},
		&models.ProductReview{})

	// Route GET cơ bản
	routes.UserRoute(r)

	protected := r.Group("/api")

	protected.Use(auth.KeycloakMiddleware())

	routes.AdminRoute(protected)
	// Chạy server trên port 8080
	r.Run(":8080")

}
