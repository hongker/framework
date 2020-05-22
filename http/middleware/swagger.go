package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hongker/framework/config"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// SwaggerHandler
func SwaggerHandler() gin.HandlerFunc {
	return ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, config.DisableSwaggerEnv)
}
