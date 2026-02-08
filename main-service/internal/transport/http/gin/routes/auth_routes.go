package routes

import (
	"github.com/gin-gonic/gin"

	handler "ptxyz/main-service/internal/transport/http/gin/handler/auth"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authHandler *handler.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}
}
