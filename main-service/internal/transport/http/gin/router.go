package gin

import (
	"github.com/gin-gonic/gin"

	"ptxyz/main-service/internal/transport/http/gin/routes"
)

func NewRouter(deps *Dependencies) *gin.Engine{
	r := gin.Default()

	// middleware (global)
	r.Use(EmptyBodyRequest())

	api := r.Group("/api")

	routes.RegisterRegisterRoutes(api, deps.RegisterHandler)
	routes.RegisterAuthRoutes(api, deps.AuthHandler)

	return r
}