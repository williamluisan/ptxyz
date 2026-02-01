package main

import (
	"fmt"
	"log"
	"ptxyz/customer-product-service/config"

	transHttpGin "ptxyz/customer-product-service/internal/transport/http/gin"

	gormMysql "ptxyz/customer-product-service/internal/infrastructure/gorm/integration/mysql"
	gormRepo "ptxyz/customer-product-service/internal/infrastructure/gorm/repository"

	userHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/user"
	userUsecase "ptxyz/customer-product-service/internal/usecase/user"
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
	dbUserRepo := gormRepo.New(db)

	/* usecases */
	userService := userUsecase.New(dbUserRepo)

	/* transport handler */
	userHandler := userHandler.New(userService)

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		UserHandler: userHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":8003")
}