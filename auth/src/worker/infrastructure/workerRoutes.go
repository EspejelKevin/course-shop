package infrastructure

import (
	"auth/src/shared/infrastructure"
	"auth/src/shared/infrastructure/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

var settings = infrastructure.NewSettings()
var namespace = settings.Namespace
var apiVersion = settings.APIVersion
var prefix = fmt.Sprintf("/%s/%s", namespace, apiVersion)

func Routes(route *gin.Engine) {
	authGroup := route.Group(prefix)
	{
		authGroup.GET("/readiness", Readiness)
		authGroup.POST("/login", Login)
		authGroup.POST("/signin", middlewares.ValidatePayloadJSON, SignIn)
	}
}
