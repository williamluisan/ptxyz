package routes

import (
	handler "ptxyz/main-service/internal/transport/http/gin/handler/konsumentenorlimit"

	"github.com/gin-gonic/gin"
)

func RegisterKonsumerTenorLimitReoutes(rg *gin.RouterGroup, ktlProxyHandler *handler.KonsumenTenorLimitProxyHandler) {
	user := rg.Group("/konsumentenorlimit")
	{
		user.GET("/:public_id", ktlProxyHandler.GetKTLByPublicId)
	}
}