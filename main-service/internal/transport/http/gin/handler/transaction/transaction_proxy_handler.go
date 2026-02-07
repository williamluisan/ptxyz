package transaction

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

type TransactionProxyHandler struct {
	baseURL string
	httpClient	*http.Client
}

func NewTransactionProxyHandler() *TransactionProxyHandler {
	return &TransactionProxyHandler{
		baseURL: viper.GetString("TRANSACTION_SERVICE_BASE_URL"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *TransactionProxyHandler) Create(c *gin.Context) {
	url := fmt.Sprintf(
		"%s%s",
		h.baseURL,
		viper.GetString("TRANSACTION_EP"),
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
		http.MethodPost,
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
			Message: "transaction service unavailable",
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