package gin

import (
	authHandler "ptxyz/main-service/internal/transport/http/gin/handler/auth"
	konsumenTenorLimitHandler "ptxyz/main-service/internal/transport/http/gin/handler/konsumentenorlimit"
	productHandler "ptxyz/main-service/internal/transport/http/gin/handler/product"
	registerHandler "ptxyz/main-service/internal/transport/http/gin/handler/register"
	transactionHandler "ptxyz/main-service/internal/transport/http/gin/handler/transaction"
)	

type Dependencies struct {
	AuthHandler *authHandler.AuthHandler
	RegisterHandler *registerHandler.RegisterProxyHandler
	ProductProxyHandler *productHandler.ProductProxyHandler
	KonsumenTenorLimitProxyHandler *konsumenTenorLimitHandler.KonsumenTenorLimitProxyHandler
	TransactionProxyHandler *transactionHandler.TransactionProxyHandler
}