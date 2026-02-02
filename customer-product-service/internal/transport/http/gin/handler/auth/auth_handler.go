package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	entity "ptxyz/customer-product-service/internal/domain/entity/auth"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	result, err := h.authService.VerifyCredential(c.Request.Context(),
		entity.VerifyInput{
			Nik:    req.Nik,
			Password: req.Password,
		},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, VerifyResponse{
		UserPublicId: result.UserPublicId,
	})
}
