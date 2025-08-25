package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// lấy tất cả user
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// tạo user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// lấy ra id user
func GetUserIndex(c *gin.Context) {
	// Get id of url
	id := c.Param("id")
	//get the user
	var user []models.User
	config.DB.Find(&user, id)
	//respond with them
	c.JSON(200, gin.H{
		"user": user,
	})
}

// cập nhật user
func UpdateUser(c *gin.Context) {
	// Lấy ID từ URL params
	id := c.Param("id")
	//get the data of req body
	var user models.User

	// Tìm user theo ID
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Nhận dữ liệu mới từ request body
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//update it
	if err := config.DB.Model(&user).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

//func DeleteUser(c *gin.Context)  {
//	var user models.User.id
//}
