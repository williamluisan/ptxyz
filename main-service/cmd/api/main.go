package main

import (
	"log"
	"ptxyz/main-service/config"
	transHttpGin "ptxyz/main-service/internal/transport/http/gin"
	"time"

	authHandler "ptxyz/main-service/internal/transport/http/gin/handler/auth"
	authUsecase "ptxyz/main-service/internal/usecase/auth"

	internalService "ptxyz/main-service/internal/infrastructure/http/client"

	konsumenTenorLimitHandler "ptxyz/main-service/internal/transport/http/gin/handler/konsumentenorlimit"
	productHandler "ptxyz/main-service/internal/transport/http/gin/handler/product"
	registerHandler "ptxyz/main-service/internal/transport/http/gin/handler/register"

	"github.com/spf13/viper"
)

func main() {
	_, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	
	/* infrastructure */
	customerProductServiceAuth := internalService.NewCustomerProductServiceAuth(viper.GetString("CUSTOMER_PRODUCT_SERVICE_BASE_URL"))

	/* usecases */
	tokenService := authUsecase.NewTokenService(viper.GetString("JWT_SECRET"), time.Duration(viper.GetInt("JWT_TTL")))
	loginService := authUsecase.NewLoginService(customerProductServiceAuth, tokenService)

	/* transport handler */
	authHandler := authHandler.NewAuthHandler(loginService)
	registerHandler := registerHandler.NewRegisterProxyHandler(viper.GetString("CUSTOMER_PRODUCT_SERVICE_BASE_URL"))
	productHandler := productHandler.NewProductProxyHandler()
	konsumenTenorLimitHandler := konsumenTenorLimitHandler.NewKonsumenTenorLimitProxyHandler()

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		RegisterHandler: registerHandler,
		AuthHandler: authHandler,
		ProductHandler: productHandler,
		KonsumenTenorLimitHandler: konsumenTenorLimitHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":8002")
}