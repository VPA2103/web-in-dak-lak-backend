package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost user=vpa password=123456 dbname=go_db port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
	fmt.Println("Database connected successfully")
}
