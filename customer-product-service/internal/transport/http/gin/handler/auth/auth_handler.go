package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	entity "ptxyz/customer-product-service/internal/domain/entity/auth"
	"ptxyz/customer-product-service/internal/transport/http/gin/handler"
	service "ptxyz/customer-product-service/internal/usecase/auth"
)

type AuthHandler struct {
	authService service.AuthService
}

func New(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) VerifyCredential(c *gin.Context) {
	var req VerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Error: &handler.APIError{
				Code: "validation_error",
				Message: "Validation error",
				Details: handler.ParseValidationError(err),
			},
		},)
		return
	}

	result, err := h.authService.VerifyCredential(c.Request.Context(),
		entity.VerifyInput{
			Nik:    req.Nik,
			Password: req.Password,
		},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, handler.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, handler.APIResponse{
		Success: true,
		Message: "",
		Data: VerifyResponse{
			UserPublicId: result.UserPublicId,
		},
	})
}
