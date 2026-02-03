package main

import (
	"fmt"
	"log"
	"ptxyz/customer-product-service/config"

	transHttpGin "ptxyz/customer-product-service/internal/transport/http/gin"

	gormMysql "ptxyz/customer-product-service/internal/infrastructure/gorm/integration/mysql"
	gormRepo "ptxyz/customer-product-service/internal/infrastructure/gorm/repository"

	authUsecase "ptxyz/customer-product-service/internal/usecase/auth"
	userUsecase "ptxyz/customer-product-service/internal/usecase/user"

	authHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/auth"
	userHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/user"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var dsn string
	if (cfg.DB.Driver == "mysql") {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
			cfg.DB.User, cfg.DB.Pass, cfg.DB.Driver, cfg.DB.Port, cfg.DB.Name, cfg.DB.Charset,
		)
	}
	db, err := gormMysql.NewMysqlDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	/* infrastructure */
	dbUserRepo := gormRepo.NewUserRepository(db)

	/* usecases */
	authService := authUsecase.New(dbUserRepo)
	userService := userUsecase.New(dbUserRepo)

	/* transport handler */
	authHandler := authHandler.New(authService)
	userHandler := userHandler.New(userService)

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":8003")
}