package gin

import (
	userHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/user"
)	

type Dependencies struct {
	UserHandler *userHandler.UserHandler
}