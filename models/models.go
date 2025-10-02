package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null;size:100" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Email     string    `gorm:"unique;not null;size:150" json:"email"`
	FullName  string    `gorm:"size:150" json:"fullName"`
	Role      string    `gorm:"type:varchar(20);default:'USR'" json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;size:200" json:"name"`
	Slug        string    `gorm:"unique;not null;size:200" json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Product struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	Name            string          `form:"name" binding:"required" json:"name"`
	Code            string          `form:"code" binding:"required" json:"code"`
	Slug            string          `form:"slug" binding:"required" json:"slug"`
	Description     string          `form:"description" json:"description"`
	Price           float64         `form:"price" json:"price"`
	Views           int             `form:"views" json:"views"`
	Rating          float32         `form:"rating" json:"rating"`
	MetaTitle       string          `form:"meta_title" json:"metaTitle"`
	MetaDescription string          `form:"meta_description" json:"metaDescription"`
	MetaKeywords    string          `form:"meta_keywords" json:"metaKeywords"`
	OgImage         string          `form:"ogImage" json:"ogImage"`
	CategoryID      uint            `form:"categoryId" json:"categoryId"`
	BrandID         uint            `form:"brandId" json:"brandId"`
	ProductImages   []ProductImages `gorm:"foreignKey:ProductID" json:"images"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	Category        Category        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
}

type ProductImages struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `json:"productId"`
	ImageURL  string  `json:"imageUrl"`
	Product   Product `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}

//type ProductSpec struct {
//	ID        uint    `gorm:"primaryKey" json:"id"`
//	ProductID uint    `json:"productId"`
//	SpecName  string  `gorm:"size:150" json:"specName"`
//	SpecValue string  `gorm:"size:255" json:"specValue"`
//	Product   Product `gorm:"constraint:OnDelete:CASCADE;" json:"product"`
//}

type ProductReview struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `json:"productId"`
	UserName  string    `gorm:"size:150;not null" json:"userName"`
	Email     string    `gorm:"size:150" json:"email"`
	Rating    int       `gorm:"check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	Product   Product   `gorm:"constraint:OnDelete:CASCADE;" json:"product"`
}

type News struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Title           string    `gorm:"not null;size:200" json:"title"`
	Slug            string    `gorm:"unique;not null;size:200" json:"slug"`
	Content         string    `gorm:"not null" json:"content"`
	Image           string    `json:"image"`
	MetaTitle       string    `json:"metaTitle"`
	MetaDescription string    `json:"metaDescription"`
	MetaKeywords    string    `json:"metaKeywords"`
	OgImage         string    `json:"ogImage"`
	AuthorID        uint      `json:"authorId"`
	Author          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type Contact struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:150;not null" json:"name"`
	Email     string    `gorm:"size:150;not null" json:"email"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Message   string    `gorm:"not null" json:"message"`
	Status    string    `gorm:"type:varchar(20);default:'new'" json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type Brand struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;size:150" json:"name"`
	Slug      string    `gorm:"unique;not null;size:150" json:"slug"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
	Products  []Product `gorm:"foreignKey:BrandID" json:"products,omitempty"`
}

//type Slide struct {
//	ID        uint      `gorm:"primaryKey" json:"id"`
//	ImageURL  string    `gorm:"not null" json:"imageUrl"`
//	Title     string    `json:"title"`
//	Link      string    `json:"link"`
//	CreatedAt time.Time `json:"createdAt"`
//}

//type Cart struct {
//	ID        uint      `gorm:"primaryKey" json:"id"`
//	UserID    uint      `json:"userId"`
//	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
//	CreatedAt time.Time `json:"createdAt"`
//}

//type CartItem struct {
//	ID        uint    `gorm:"primaryKey" json:"id"`
//	CartID    uint    `json:"cartId"`
//	ProductID uint    `json:"productId"`
//	Quantity  int     `gorm:"default:1" json:"quantity"`
//	Price     float64 `gorm:"type:decimal(12,2)" json:"price"`
//	Cart      Cart    `gorm:"constraint:OnDelete:CASCADE;" json:"cart"`
//	Product   Product `json:"product"`
//}

//type Order struct {
//	ID          uint      `gorm:"primaryKey" json:"id"`
//	UserID      uint      `json:"userId"`
//	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
//	TotalAmount float64   `gorm:"type:decimal(12,2)" json:"totalAmount"`
//	Status      string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
//	CreatedAt   time.Time `json:"createdAt"`
//}

//type OrderItem struct {
//	ID        uint    `gorm:"primaryKey" json:"id"`
//	OrderID   uint    `json:"orderId"`
//	ProductID uint    `json:"productId"`
//	Quantity  int     `json:"quantity"`
//	Price     float64 `gorm:"type:decimal(12,2)" json:"price"`
//	Order     Order   `gorm:"constraint:OnDelete:CASCADE;" json:"order"`
//	Product   Product `json:"product"`
//}

type UserAddress struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"userId"`
	Name      string `gorm:"size:150;not null" json:"name"`
	Phone     string `gorm:"size:20;not null" json:"phone"`
	Address   string `gorm:"not null" json:"address"`
	IsDefault bool   `gorm:"default:false" json:"isDefault"`
	User      User   `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
}
