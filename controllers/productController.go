package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind dữ liệu form (trừ ảnh)
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lấy nhiều ảnh từ form-data (key = "images")
	form, _ := c.MultipartForm()
	files := form.File["images"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images uploaded"})
		return
	}

	// Lưu product trước (để có ID)
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product"})
		return
	}

	var imageURLs []string

	for i, file := range files {
		src, err := file.Open()
		if err != nil {
			continue
		}

		// Upload Cloudinary
		resp, err := config.CLD.Upload.Upload(c, src, uploader.UploadParams{})
		src.Close()
		if err != nil {
			continue // bỏ qua file lỗi
		}

		// Lưu vào bảng product_images
		productImage := models.ProductImages{
			ProductID: product.ID,
			ImageURL:  resp.SecureURL,
		}
		config.DB.Create(&productImage)

		imageURLs = append(imageURLs, resp.SecureURL)

		// Ảnh đầu tiên sẽ làm OgImage (ảnh đại diện)
		if i == 0 {
			product.OgImage = resp.SecureURL
			config.DB.Model(&product).Update("og_image", resp.SecureURL)
		}
	}
	config.DB.Preload("ProductImages").First(&product, product.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo sản phẩm thành công",
		"product": product,
		"images":  imageURLs,
	})
}

//

func GetProduct(c *gin.Context) {
	var products []models.Product
	config.DB.Preload("Category").Preload("ProductImages").Find(&products)

	c.JSON(http.StatusOK, products)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	// Kiểm tra sản phẩm tồn tại không
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Xóa ảnh liên quan trước (nếu có)
	if err := config.DB.Where("product_id = ?", id).Delete(&models.ProductImages{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product images"})
		return
	}

	// Xóa sản phẩm
	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func GetProductIndex(c *gin.Context) {
	// Get id of url
	id := c.Param("id")
	//get the user
	var product models.Product
	if err := config.DB.Preload("Category").Preload("ProductImages").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	//respond with them
	c.JSON(200, gin.H{
		"product": product,
	})
}
