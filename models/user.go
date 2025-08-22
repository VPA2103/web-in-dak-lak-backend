package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // tự động ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name"`
}
