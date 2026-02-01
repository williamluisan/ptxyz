package routes

import (
	"github.com/gin-gonic/gin"

	handler "ptxyz/main-service/internal/transport/http/gin/handler/register"
)

func RegisterRegisterRoutes(rg *gin.RouterGroup, registerHandler *handler.RegisterProxyHandler) {
	rg.POST("/register", registerHandler.Register)
}
