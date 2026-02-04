package transaction

import (
	"context"
	"net/http"
	entity "ptxyz/transaction-service/internal/domain/entity/transaction"
	"ptxyz/transaction-service/internal/transport/http/gin/handler"
	service "ptxyz/transaction-service/internal/usecase/transaction"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var req CreateTransactionRequest

	token := c.GetHeader("Authorization")

	ctx := context.WithValue(
		c.Request.Context(),
		"authorization",
		token,
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Error: &handler.APIError{
				Code: "validation_error",
				Message: "Validation error",
				Details: handler.ParseValidationError(err),
			},
		})
		return
	}

	input := &entity.TransactionInput{
		KonsumenTenorLimitPublicId: req.KonsumenTenorLimitPublicId,	
		ProductPublicId: req.ProductPublicId,				
		NomorKontrak: req.NomorKontrak,				
		HargaOtr: req.HargaOtr,					
		AdminFee: req.AdminFee,					
		JumlahCicilan: req.JumlahCicilan,				
		JumlahBunga: req.JumlahBunga,					
		NamaAset: req.NamaAset,					
	}

	if err := h.transactionService.Create(ctx, input); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, handler.APIResponse{
		Success: true,
		Message: "Berhasil melakukan transaksi",
	})
}