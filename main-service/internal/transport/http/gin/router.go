package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"ptxyz/main-service/internal/transport/http/gin/routes"

	middleware "ptxyz/main-service/internal/transport/http/gin/middleware"
)

func NewRouter(deps *Dependencies) *gin.Engine{
	r := gin.Default()

	// middleware (global)
	r.Use(middleware.EmptyBodyRequest())

	api := r.Group("/api")

	routes.RegisterAuthRoutes(api, deps.AuthHandler)
	api.Use(middleware.JWT(viper.GetString("JWT_SECRET")))
	{
		routes.RegisterRegisterRoutes(api, deps.RegisterHandler)
	}

	return r
}