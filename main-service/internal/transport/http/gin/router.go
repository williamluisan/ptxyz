package gin

import (
	"github.com/gin-gonic/gin"

	"ptxyz/main-service/internal/transport/http/gin/routes"

	middleware "ptxyz/main-service/internal/transport/http/gin/middleware"
)

func NewRouter(deps *Dependencies) *gin.Engine{
	r := gin.Default()

	// middleware (global)
	r.Use(middleware.EmptyBodyRequest())

	api := r.Group("/api")

	routes.RegisterAuthRoutes(api, deps.AuthHandler)
	routes.RegisterRegisterRoutes(api, deps.RegisterHandler)
	routes.RegisterProductRoutes(api, deps.ProductHandler)
	routes.RegisterKonsumerTenorLimitReoutes(api, deps.KonsumenTenorLimitHandler)

	return r
}