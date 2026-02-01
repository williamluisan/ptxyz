package user

import (
	"fmt"
	"net/http"
	"time"

	entity "ptxyz/customer-product-service/internal/domain/entity/user"
	"ptxyz/customer-product-service/internal/infrastructure/storage"
	"ptxyz/customer-product-service/internal/transport/http/gin/handler"
	service "ptxyz/customer-product-service/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type UserHandler struct {
	userService service.UserService
	storage storage.LocalStorage
}

func New(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetByNik(c *gin.Context) {
	c.JSON(http.StatusOK, "user@gmail.com")
}

func (h *UserHandler) Create(c *gin.Context) {
	var req CreateUserRequest

	ctx := c

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
	
	/* file upload */
	h.storage.AllowedExt = []string{".jpg", ".png"}

	fileKtpPath := fmt.Sprintf("users/%s/%s", req.Nik, req.FotoKtp.Filename)
	_, err := h.storage.Save(ctx, fileKtpPath, req.FotoKtp)
	if err != nil {
		c.JSON(http.StatusBadGateway, handler.APIResponse{
			Success: false,
			Message: "Failed to upload KTP file.",
		})
		return
	}
	
	fileSelfiePath := fmt.Sprintf("users/%s/%s", req.Nik, req.FotoSelfie.Filename)
	_, err = h.storage.Save(ctx, fileSelfiePath, req.FotoSelfie)
	if err != nil {
		c.JSON(http.StatusBadGateway, handler.APIResponse{
			Success: false,
			Message: "Failed to upload KTP file.",
		})
		return
	}
	/* // */

	/* date and decimal parsing */
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
	/* // */

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