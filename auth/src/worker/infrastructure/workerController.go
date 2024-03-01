package infrastructure

import (
	"auth/src/shared/domain"
	"auth/src/worker/container"

	"github.com/gin-gonic/gin"
)

func Readiness(ctx *gin.Context) {
	usecase := container.ContainerReadiness()
	data := usecase.Execute(ctx)
	switch content := data.(type) {
	case domain.FailureResponse:
		ctx.JSON(content.StatusCode, content.Response)
	case domain.SuccessResponse:
		ctx.JSON(content.StatusCode, content.Response)
	}
}

func Login(ctx *gin.Context) {
	usecase := container.ContainerLogIn()
	data := usecase.Execute(ctx)
	switch content := data.(type) {
	case domain.FailureResponse:
		ctx.JSON(content.StatusCode, content.Response)
	case domain.SuccessResponse:
		ctx.JSON(content.StatusCode, content.Response)
	}
}

func SignIn(ctx *gin.Context) {
	usecase := container.ContainerSignIn()
	data := usecase.Execute(ctx)
	switch content := data.(type) {
	case domain.FailureResponse:
		ctx.JSON(content.StatusCode, content.Response)
	case domain.SuccessResponse:
		ctx.JSON(content.StatusCode, content.Response)
	}
}
