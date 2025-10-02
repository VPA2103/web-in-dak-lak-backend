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
	//1.
	//err := auth.InitKeycloak()
	//if err != nil {
	//	panic(err)
	//}

	config.InitCloudinary()

	//2. Tạo router mặc định
	r := gin.Default()
	config.SetupCORS(r)

	//
	config.ConnectDB()
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Brand{},
		&models.Product{},
		&models.ProductImages{},
		&models.ProductReview{},
		&models.News{},
		&models.Contact{},
		&models.UserAddress{},
	); err != nil {
		panic("Migration failed: " + err.Error())
	}

	// Route GET cơ bản
	routes.UserRoute(r)
	routes.ProductRoute(r)
	routes.CategoryRoute(r)

	//protected := r.Group("/api")

	//protected.Use(auth.KeycloakMiddleware())

	routes.UploadRoutes(r)

	//routes.AdminRoute(protected)
	// Chạy server trên port 8080
	r.Run(":8080")

}
