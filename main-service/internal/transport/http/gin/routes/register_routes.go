package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	handler "ptxyz/main-service/internal/transport/http/gin/handler/register"
	middleware "ptxyz/main-service/internal/transport/http/gin/middleware"
)

func RegisterRegisterRoutes(rg *gin.RouterGroup, registerHandler *handler.RegisterProxyHandler) {
	rg.Use(middleware.JWT(viper.GetString("JWT_SECRET"))) 
	{
		rg.POST("/register", registerHandler.Register)
	}
}
