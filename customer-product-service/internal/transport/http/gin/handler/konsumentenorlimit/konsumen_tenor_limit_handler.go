package konsumentenorlimit

import (
	"fmt"
	"net/http"
	entity "ptxyz/customer-product-service/internal/domain/entity/konsumentenorlimit"
	"ptxyz/customer-product-service/internal/transport/http/gin/handler"
	service "ptxyz/customer-product-service/internal/usecase/konsumentenorlimit"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

type KonsumenTenorLimitHandler struct {
	konsumenTenorLimitService service.KonsumenTenorLimitService
}

func New(konsumenTenorLimitService service.KonsumenTenorLimitService) *KonsumenTenorLimitHandler {
	return &KonsumenTenorLimitHandler{konsumenTenorLimitService: konsumenTenorLimitService}
}

func (h *KonsumenTenorLimitHandler) GetByPublicId(c *gin.Context) {
	publicId := c.Param("public_id")

	if _, err := ulid.Parse(publicId); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Error: &handler.APIError{
				Code: "wrong_parameter_format",
				Message: "Wrong parameter format",
			},
		})
		return
	}

	data, err := h.konsumenTenorLimitService.GetByPublicId(c, publicId)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, handler.APIResponse{
		Success: true,
		Data: handler.APIResponse{
			Success: true,
			Data: data,
		},
	})
}

func (h *KonsumenTenorLimitHandler) UpdateBalance(c *gin.Context) {
	var req UpdateKTLBalanceRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
				Success: false,
				Error: &handler.APIError{
					Code: "validation_error",
					Message: "Validation error",
					Details: handler.ParseValidationError(err),
				},
			},
		)
		return
	}

	update := &entity.KTLUpdateBalance{
		PublicId: req.PublicId,
		UsedAmount: req.UsedAmount,
		Balance: req.Balance,
		UpdatedBy: 1,
	}

	if err := h.konsumenTenorLimitService.UpdateBalance(c, update); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, handler.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Berhasil update balance konsumer tenor limit. Public Id: %s", req.PublicId),
	})
}