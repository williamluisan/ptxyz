package gin

import (
	authHandler "ptxyz/main-service/internal/transport/http/gin/handler/auth"
	registerHandler "ptxyz/main-service/internal/transport/http/gin/handler/register"
)	

type Dependencies struct {
	AuthHandler *authHandler.AuthHandler
	RegisterHandler *registerHandler.RegisterProxyHandler
}