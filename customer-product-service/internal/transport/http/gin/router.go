package gin

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"ptxyz/customer-product-service/internal/transport/http/gin/routes"
)

func NewRouter(deps *Dependencies) *gin.Engine{
	// register tag name function (before route initialization)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	// initialize the route
	r := gin.Default()

	// middleware (global)
	r.Use(EmptyBodyRequest())

	api := r.Group("/api")

	routes.RegisterUserRoutes(api, deps.UserHandler)

	return r
}