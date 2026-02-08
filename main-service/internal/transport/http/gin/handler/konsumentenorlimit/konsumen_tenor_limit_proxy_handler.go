package konsumentenorlimit

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"ptxyz/main-service/internal/transport/http/gin/handler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type KonsumenTenorLimitProxyHandler struct {
	baseURL string
	httpClient	*http.Client
}

func NewKonsumenTenorLimitProxyHandler() *KonsumenTenorLimitProxyHandler {
	return &KonsumenTenorLimitProxyHandler{
		baseURL: viper.GetString("CUSTOMER_PRODUCT_SERVICE_BASE_URL"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *KonsumenTenorLimitProxyHandler) GetKTLByPublicId(c *gin.Context) {
	url := fmt.Sprintf(
		"%s%s/%s",
		h.baseURL,
		viper.GetString("CUSTOMER_PRODUCT_KONSUMEN_TENOR_LIMIT_EP"),
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
			Message: "konsumer tenor limit service unavailable",
		})
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

func (h *KonsumenTenorLimitProxyHandler) UpdateBalance(c *gin.Context) {
	url := fmt.Sprintf(
		"%s%s",
		h.baseURL,
		viper.GetString("CUSTOMER_PRODUCT_KONSUMEN_TENOR_LIMIT_UPDATE_BALANCE_EP"),
	)

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, handler.APIResponse{
			Success: false,
			Message: "failed to read request body",
		})
		return
	}

	req, err := http.NewRequestWithContext(
		c.Request.Context(),
		http.MethodPut,
		url,
		bytes.NewBuffer(bodyBytes),
	)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, handler.APIResponse{
			Success: false,
			Message: "Failed to create proxy request",
		})
		return
	}

	req.Header = c.Request.Header.Clone()
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, handler.APIResponse{
			Success: false,
			Message: "costumer product service unavailable",
		})
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}