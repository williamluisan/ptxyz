package gin

import (
	authHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/auth"
	userHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/user"
)	

type Dependencies struct {
	UserHandler *userHandler.UserHandler
	AuthHandler *authHandler.AuthHandler
}