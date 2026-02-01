package gin

import (
	"github.com/gin-gonic/gin"

	"ptxyz/main-service/internal/transport/http/gin/routes"
)

func NewRouter(deps *Dependencies) *gin.Engine{
	r := gin.Default()

	api := r.Group("/api")

	routes.RegisterRegisterRoutes(api, deps.RegisterHandler)

	return r
}