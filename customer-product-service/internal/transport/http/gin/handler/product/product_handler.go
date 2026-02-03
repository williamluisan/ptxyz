package product

import (
	"net/http"
	"ptxyz/customer-product-service/internal/transport/http/gin/handler"
	productUsecase "ptxyz/customer-product-service/internal/usecase/product"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

type ProductHandler struct {
	productService	productUsecase.ProductService
}

func NewProductHandler(productService productUsecase.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetByPublicId(c *gin.Context) {
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

	data, err := h.productService.GetByPublicId(c, publicId)
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