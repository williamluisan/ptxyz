package routes

import (
	handler "ptxyz/transaction-service/internal/transport/http/gin/handler/transaction"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(rg *gin.RouterGroup, h *handler.TransactionHandler) {
	order := rg.Group("/transaction")
	{
		order.POST("/", h.Create)
	}
}