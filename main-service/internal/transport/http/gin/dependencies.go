package gin

import (
	registerHandler "ptxyz/main-service/internal/transport/http/gin/handler/register"
)	

type Dependencies struct {
	RegisterHandler *registerHandler.RegisterProxyHandler
}