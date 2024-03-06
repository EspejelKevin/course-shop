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
var prefix = fmt.Sprintf("/%s/api/%s", namespace, apiVersion)

func Routes(route *gin.Engine) {
	healthChecks := route.Group(prefix)
	{
		healthChecks.GET("/liveness", Liveness)
		healthChecks.GET("/readiness", Readiness)
	}

	authGroup := route.Group(prefix)
	{
		authGroup.POST("/login", middlewares.ValidatePayloadLogIn, Login)
		authGroup.POST("/signup", middlewares.ValidatePayloadSignIn, SignUp)
	}

	validationsGroup := route.Group(prefix)
	{
		validationsGroup.GET("/validations/token", middlewares.ValidateBearerToken, ValidateToken)
		validationsGroup.POST("/validations/email", middlewares.ValidateVerificationCode, ValidateEmail)
		validationsGroup.POST("/validations/phone", middlewares.ValidateBearerToken)
	}

	confirmationsGroup := route.Group(prefix)
	{
		confirmationsGroup.POST("/confirmations/phone", middlewares.ValidateBearerToken, ConfirmPhone)
		confirmationsGroup.POST("/confirmations/email", middlewares.ValidateEmailData, ConfirmEmail)
	}
}
