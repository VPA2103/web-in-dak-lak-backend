package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoute(rg *gin.RouterGroup) {

	rg.GET("/admin", func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to admin area - protected by Keycloak!",
			"user":    user,
		})
	})

}
