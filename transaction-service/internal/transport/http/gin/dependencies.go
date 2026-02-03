package gin

import (
	transactionHandler "ptxyz/transaction-service/internal/transport/http/gin/handler/transaction"
)	

type Dependencies struct {
	TransactionHandler *transactionHandler.TransactionHandler
}