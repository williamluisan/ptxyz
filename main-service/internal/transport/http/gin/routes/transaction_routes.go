package routes

import (
	"github.com/gin-gonic/gin"

	transactionProxyHandler "ptxyz/main-service/internal/transport/http/gin/handler/transaction"
)

func RegisterTransactionRoutes(rg *gin.RouterGroup, transactionProxyHandler *transactionProxyHandler.TransactionProxyHandler) {
	user := rg.Group("/transaction")
	{
		user.POST("/", transactionProxyHandler.Create)
	}
}
