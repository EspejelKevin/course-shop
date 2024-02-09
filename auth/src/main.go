package main

import (
	"auth/src/shared/infrastructure"
	workerInfrastructure "auth/src/worker/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	settings := infrastructure.NewSettings()
	router := gin.Default()
	workerInfrastructure.Routes(router)
	router.Run(settings.Port)
}
