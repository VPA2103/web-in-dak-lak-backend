package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func KeycloakMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		idToken, err := Verifier.Verify(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Gắn claims vào context để route sử dụng
		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err == nil {
			c.Set("user", claims)
		}

		c.Next()
	}
}
