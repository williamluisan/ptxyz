package gin

import (
	"net/http"
	"ptxyz/customer-product-service/internal/transport/http/gin/handler"

	"github.com/gin-gonic/gin"
)

func EmptyBodyRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodPatch:
			if c.Request.ContentLength == 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, handler.APIResponse{
					Success: false,
					Error: &handler.APIError{
						Code: "empty_body",
						Message: "Empty Body",
						Details: "Empty Body",
					},
				})
				return
			}
		}
	}
}