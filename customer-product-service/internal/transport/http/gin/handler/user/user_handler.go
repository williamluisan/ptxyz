package user

import (
	"net/http"

	entity "ptxyz/customer-product-service/internal/domain/entity/user"
	"ptxyz/customer-product-service/internal/transport/http/gin/handler"
	service "ptxyz/customer-product-service/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetByNik(c *gin.Context) {
	c.JSON(http.StatusOK, "user@gmail.com")
}

func (h *UserHandler) Create(c *gin.Context) {
	var req CreateUserRequest

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
	
	input := &entity.UserInput{
		Nik: req.Nik,
		FullName: req.FullName,
		LegalName: req.LegalName,
		TempatLahir: req.TempatLahir,
		TanggalLahir: req.TanggalLahir,
		Gaji: req.Gaji,
		FotoKtp: req.FotoKtp.Filename, 
		FotoSelfie: req.FotoSelfie.Filename,
	}

	if err := h.userService.Create(c, input); err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
				Success: false,
				Message: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, handler.APIResponse{
		Success: true,
		Message: "User created successfully",
	})
}