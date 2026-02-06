package routes

import (
	handler "ptxyz/main-service/internal/transport/http/gin/handler/product"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, productHandler *handler.ProductProxyHandler) {
	rg.GET("/product/:public_id", productHandler.GetProductByPublicId)
}