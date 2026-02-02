package routes

import (
	userHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *userHandler.UserHandler) {
	user := rg.Group("/user")
	{
		user.POST("/register", h.Create)
		user.GET("/nik/:nik", h.GetByNik)
	}
}