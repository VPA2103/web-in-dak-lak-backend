package controllers

import (
	"backend/config"
	"backend/models"

	//"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	// Lấy file từ request
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Mở file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	// Upload lên Cloudinary
	resp, err := config.CLD.Upload.Upload(c, src, uploader.UploadParams{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload success",
		"url":     resp.SecureURL,
	})
}

func GetProductImage(c *gin.Context) {
	var images []models.ProductImages
	if err := config.DB.Find(&images).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product images"})
		return
	}
	c.JSON(http.StatusOK, images)
}
