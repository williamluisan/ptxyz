package gin

import (
	authHandler "ptxyz/main-service/internal/transport/http/gin/handler/auth"
	konsumenTenorLimitHandler "ptxyz/main-service/internal/transport/http/gin/handler/konsumentenorlimit"
	productHandler "ptxyz/main-service/internal/transport/http/gin/handler/product"
	registerHandler "ptxyz/main-service/internal/transport/http/gin/handler/register"
)	

type Dependencies struct {
	AuthHandler *authHandler.AuthHandler
	RegisterHandler *registerHandler.RegisterProxyHandler
	ProductHandler *productHandler.ProductProxyHandler
	KonsumenTenorLimitHandler *konsumenTenorLimitHandler.KonsumenTenorLimitProxyHandler
}