package main

import (
	"fmt"
	"log"
	"ptxyz/transaction-service/config"

	transHttpGin "ptxyz/transaction-service/internal/transport/http/gin"

	gormMysql "ptxyz/transaction-service/internal/infrastructure/gorm/integration/mysql"
	gormRepo "ptxyz/transaction-service/internal/infrastructure/gorm/repository"

	transactionUsecase "ptxyz/transaction-service/internal/usecase/transaction"

	transactionHandler "ptxyz/transaction-service/internal/transport/http/gin/handler/transaction"

	internalService "ptxyz/transaction-service/internal/infrastructure/http/client"

	"github.com/spf13/viper"
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
	dbTransactionRepo := gormRepo.NewTransactionRepository(db)
	konsumenTenorLimitService := internalService.NewKonsumenTenorLimit(viper.GetString("CUSTOMER_PRODUCT_SERVICE_BASE_URL"))

	/* usecases */
	transactionService := transactionUsecase.NewTransactionService(dbTransactionRepo, konsumenTenorLimitService)

	/* transport handler */
	transactionHandler := transactionHandler.NewTransactionHandler(transactionService)

	/* transport dependencies */
	deps := &transHttpGin.Dependencies{
		TransactionHandler: transactionHandler,
	}

	// router
	router := transHttpGin.NewRouter(deps)
	router.Run(":8004")
}

