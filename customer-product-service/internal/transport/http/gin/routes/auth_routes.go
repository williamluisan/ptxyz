package routes

import (
	authHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, h *authHandler.AuthHandler) {
	user := rg.Group("/auth")
	{
		user.POST("/login", h.VerifyCredential)
	}
}