package main

import (
	"auth/src/shared/infrastructure"
	workerInfrastructure "auth/src/worker/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	settings := infrastructure.NewSettings()
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.Use(gin.CustomRecovery(infrastructure.InternalServerErrorHandler))
	router.NoRoute(infrastructure.NotFoundErrorHandler)
	router.NoMethod(infrastructure.MethodNotAllowedErrorHandler)
	workerInfrastructure.Routes(router)
	router.Run(settings.Port)
}
