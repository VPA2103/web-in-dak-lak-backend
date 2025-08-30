package models

import (
	"time"
)

// Bảng users
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"size:100;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:150;unique;not null"`
	FullName  string    `gorm:"size:150"`
	Role      string    `gorm:"type:varchar(20);default:'user'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// Bảng categories (bạn tham chiếu category_id nên mình thêm)
type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:150;not null"`
	Slug string `gorm:"size:200;unique;not null"`
}

// Bảng products
type Product struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	Name            string    `gorm:"size:200;not null"`
	Code            string    `gorm:"size:100;unique;not null"`
	Slug            string    `gorm:"size:200;unique;not null"`
	Description     string    `gorm:"type:text"`
	Price           float64   `gorm:"type:decimal(12,2)"`
	Views           int       `gorm:"default:0"`
	Rating          float32   `gorm:"type:decimal(2,1);default:0"`
	MetaTitle       string    `gorm:"size:255"`
	MetaDescription string    `gorm:"size:300"`
	MetaKeywords    string    `gorm:"size:255"`
	OgImage         string    `gorm:"size:255"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	CategoryID      uint
	Category        Category `gorm:"foreignKey:CategoryID"`
}

// Bảng news
type News struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	Title           string `gorm:"size:200;not null"`
	Slug            string `gorm:"size:200;unique;not null"`
	Content         string `gorm:"type:text;not null"`
	Image           string `gorm:"size:255"`
	MetaTitle       string `gorm:"size:255"`
	MetaDescription string `gorm:"size:300"`
	MetaKeywords    string `gorm:"size:255"`
	OgImage         string `gorm:"size:255"`
	AuthorID        uint
	Author          User      `gorm:"foreignKey:AuthorID"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

// Bảng contacts
type Contact struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"size:150;not null"`
	Email   string `gorm:"size:150;not null"`
	Phone   string `gorm:"size:20"`
	Message string `gorm:"type:text;not null"`
	Status  string `gorm:"type:varchar(20);default:'new'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// Bảng slides
type Slide struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	ImageURL  string    `gorm:"size:255;not null"`
	Title     string    `gorm:"size:200"`
	Link      string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// Bảng product_images
type ProductImage struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	ProductID uint    `gorm:"not null"`
	ImageURL  string  `gorm:"size:255;not null"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

// Bảng product_specs
type ProductSpec struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	ProductID uint    `gorm:"not null"`
	SpecName  string  `gorm:"size:150;not null"`
	SpecValue string  `gorm:"size:255;not null"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

// Bảng product_reviews
type ProductReview struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	ProductID uint      `gorm:"not null"`
	UserName  string    `gorm:"size:150;not null"`
	Email     string    `gorm:"size:150"`
	Rating    int       `gorm:"check:rating >= 1 AND rating <= 5"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}
