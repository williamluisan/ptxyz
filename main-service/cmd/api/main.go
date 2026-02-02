package main

import (
	transHttpGin "ptxyz/main-service/internal/transport/http/gin"
	registerHandler "ptxyz/main-service/internal/transport/http/gin/handler/register"
)

func main() {
	registerHandler := registerHandler.NewRegisterProxyHandler("a")
	deps := &transHttpGin.Dependencies{
		RegisterHandler: registerHandler,
	}

	router := transHttpGin.NewRouter(deps)
	router.Run(":8000")
}