package routes

import (
	konsumenTenorLimitHandler "ptxyz/customer-product-service/internal/transport/http/gin/handler/konsumentenorlimit"

	"github.com/gin-gonic/gin"
)

func RegisterKonsumenTenorLimitRoutes(rg *gin.RouterGroup, h *konsumenTenorLimitHandler.KonsumenTenorLimitHandler) {
	konsumenTenorLimit := rg.Group("/konsumentenorlimit") 
	{
		konsumenTenorLimit.GET("/:public_id", h.GetByPublicId)
		konsumenTenorLimit.PUT("/updateBalance", h.UpdateBalance)
	}
}