package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RegisterProxyHandler struct {
	registerServiceBaseURL string
	httpClient             *http.Client
}

func NewRegisterProxyHandler(registerServiceBaseURL string) *RegisterProxyHandler {
	return &RegisterProxyHandler{
		registerServiceBaseURL: registerServiceBaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *RegisterProxyHandler) Register(c *gin.Context) {
	targetURL := h.registerServiceBaseURL + viper.GetString("CUSTOMER_PRODUCT_REGISTER_EP")
fmt.Println(targetURL)
	req, err := http.NewRequest(
		c.Request.Method,
		targetURL,
		c.Request.Body,
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create proxy request",
		})
		return
	}

	req.Header = c.Request.Header.Clone()

	if userID, ok := c.Get("user_public_id"); ok {
		req.Header.Set("X-User-ID", userID.(string))
	}

	req.Header.Set("X-Forwarded-For", c.ClientIP())
	req.Header.Set("X-Request-ID", c.GetString("request_id"))

	resp, err := h.httpClient.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "customer service unavailable",
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
