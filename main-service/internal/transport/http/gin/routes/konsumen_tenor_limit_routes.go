package routes

import (
	handler "ptxyz/main-service/internal/transport/http/gin/handler/konsumentenorlimit"

	"github.com/gin-gonic/gin"
)

func RegisterKonsumerTenorLimitReoutes(rg *gin.RouterGroup, ktlProxyHandler *handler.KonsumenTenorLimitProxyHandler) {
	ktl := rg.Group("/konsumentenorlimit")
	{
		ktl.GET("/:public_id", ktlProxyHandler.GetKTLByPublicId)
		ktl.PUT("/updateBalance", ktlProxyHandler.UpdateBalance)
	}
}