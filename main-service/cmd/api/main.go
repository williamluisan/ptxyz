package main

import (
	transHttpGin "ptxyz/main-service/internal/transport/http/gin"

	authHandler "ptxyz/main-service/internal/transport/http/gin/handler/auth"
	authUsecase "ptxyz/main-service/internal/usecase/auth"

	internalService "ptxyz/main-service/internal/infrastructure/http/client"

	registerHandler "ptxyz/main-service/internal/transport/http/gin/handler/register"
)

func main() {
	/* infrastructure */
	customerProductServiceAuth := internalService.NewCustomerProductServiceAuth("http://localhost:8003/api")

	/* usecases */
	tokenService := authUsecase.NewTokenService("", 100)
	loginService := authUsecase.NewLoginService(customerProductServiceAuth, tokenService)

	/* transport handler */
	authHandler := authHandler.NewAuthHandler(loginService)
	registerHandler := registerHandler.NewRegisterProxyHandler("a")

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		RegisterHandler: registerHandler,
		AuthHandler: authHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":8002")
}