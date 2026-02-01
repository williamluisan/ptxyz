package user

import (
	"net/http"
	"time"

	entity "ptxyz/customer-product-service/internal/domain/entity/user"
	"ptxyz/customer-product-service/internal/transport/http/gin/handler"
	service "ptxyz/customer-product-service/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type UserHandler struct {
	userService service.UserService
}

func New(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetByNik(c *gin.Context) {
	c.JSON(http.StatusOK, "user@gmail.com")
}

func (h *UserHandler) Create(c *gin.Context) {
	var req CreateUserRequest

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
	
	tanggalLahir, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: "Invalid tanggalLahir format, use YYYY-MM-DD",
		})
		return
	}

	gaji, err := decimal.NewFromString(req.Gaji)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: "Invalid gaji format",
		})
		return
	}

	input := &entity.UserInput{
		Nik: req.Nik,
		FullName: req.FullName,
		LegalName: req.LegalName,
		TempatLahir: req.TempatLahir,
		TanggalLahir: tanggalLahir,
		Gaji: gaji,
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