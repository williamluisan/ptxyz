package routes

import (
	handler "ptxyz/main-service/internal/transport/http/gin/handler/product"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, productHandler *handler.ProductProxyHandler) {
	product := rg.Group("/product")
	{
		product.GET("/:public_id", productHandler.GetProductByPublicId)
	}
}