package gin

import (
	authHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/auth"
	konsumenTenorLimit "ptxyz/customer-product-service/internal/transport/http/gin/handler/konsumen_tenor_limit"
	productHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/product"
	userHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/user"
)	

type Dependencies struct {
	UserHandler *userHandler.UserHandler
	AuthHandler *authHandler.AuthHandler
	ProductHandler *productHandler.ProductHandler
	KonsumenTenorLimitHandler *konsumenTenorLimit.KonsumenTenorLimitHandler
}