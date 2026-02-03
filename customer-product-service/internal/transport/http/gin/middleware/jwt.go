package gin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"ptxyz/main-service/internal/transport/http/gin/handler"
)

type Claims struct {
	UserPublicID string
	TenantID     string
	Role         string
}

func JWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Message: "missing authorization header",
			}) 
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Message: "invalid authorization header",
			})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Message: "invalid token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handler.APIResponse{
				Success: false,
				Message: "invalid token claims",
			})
			return
		}

		c.Set("user_public_id", claims["sub"])

		c.Next()
	}
}