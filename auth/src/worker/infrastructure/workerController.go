package infrastructure

import (
	"auth/src/worker/container"

	"github.com/gin-gonic/gin"
)

func Readiness(ctx *gin.Context) {
	usecase := container.ContainerReadiness()
	response := usecase.Execute(ctx)
	status := response["status"].(int)
	ctx.JSON(status, response["message"])
}

func Login(ctx *gin.Context) {

}

func Register(ctx *gin.Context) {

}
