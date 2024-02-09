package infrastructure

import (
	"auth/src/worker/container"

	"github.com/gin-gonic/gin"
)

func Readiness(ctx *gin.Context) {
	usecase := container.ContainerReadiness()
	response, statusCode := usecase.Execute(ctx)
	ctx.JSON(statusCode, response)
}

func Login(ctx *gin.Context) {

}

func Register(ctx *gin.Context) {

}
