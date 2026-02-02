package auth

import (
	"net/http"

	"ptxyz/main-service/internal/transport/http/gin/handler"

	"github.com/gin-gonic/gin"

	authEntity "ptxyz/main-service/internal/domain/entity/auth"
	authUsecase "ptxyz/main-service/internal/usecase/auth"
)

type AuthHandler struct {
	loginService authUsecase.LoginService
}

func NewAuthHandler(s authUsecase.LoginService) *AuthHandler {
	return &AuthHandler{loginService: s}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
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

	result, err := h.loginService.Login(c.Request.Context(),authEntity.LoginInput{
			Nik: req.Nik,
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
		Data: LoginResponse{
			AccessToken: result.AccessToken,
		},
	})
}
