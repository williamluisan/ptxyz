package routes

import (
	"github.com/gin-gonic/gin"

	productHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/product"
)

func RegisterProductRoutes(rg *gin.RouterGroup, h *productHandler.ProductHandler) {
	product := rg.Group("/product")
	{
		product.GET("/:public_id", h.GetByPublicId)
	}
}