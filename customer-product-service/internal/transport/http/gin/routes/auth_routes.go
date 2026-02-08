package routes

import (
	authHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, h *authHandler.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", h.VerifyCredential)
	}
}