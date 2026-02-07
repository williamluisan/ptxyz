package handler

import (
	"fmt"
	"io"
	"net/http"
	"ptxyz/main-service/internal/transport/http/gin/handler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ProductProxyHandler struct {
	baseURL string
	httpClient	*http.Client
}

func NewProductProxyHandler() *ProductProxyHandler{
	return &ProductProxyHandler{
		baseURL: viper.GetString("CUSTOMER_PRODUCT_SERVICE_BASE_URL"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *ProductProxyHandler) GetProductByPublicId(c *gin.Context) {
	url := fmt.Sprintf(
		"%s%s/%s",
		h.baseURL,
		viper.GetString("CUSTOMER_PRODUCT_PRODUCT_EP"),
		c.Param("public_id"),
	)

	req, err := http.NewRequest(
		c.Request.Method,
		url,
		nil,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, handler.APIResponse{
			Success: false,
			Message: "Failed to create proxy request",
		})
		return
	}

	req.Header = c.Request.Header.Clone()

	resp, err := h.httpClient.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, handler.APIResponse{
			Success: false,
			Message: "product service unavailable",
		})
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}