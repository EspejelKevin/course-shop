package infrastructure

import (
	"auth/src/shared/domain"
	"auth/src/worker/container"

	"github.com/gin-gonic/gin"
)

func Liveness(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "Service running",
	})
}

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

func SignUp(ctx *gin.Context) {
	usecase := container.ContainerSignUp()
	data := usecase.Execute(ctx)
	switch content := data.(type) {
	case domain.FailureResponse:
		ctx.JSON(content.StatusCode, content.Response)
	case domain.SuccessResponse:
		ctx.JSON(content.StatusCode, content.Response)
	}
}

func ValidateToken(ctx *gin.Context) {
	usecase := container.ContainerValidateToken()
	data := usecase.Execute(ctx)
	switch content := data.(type) {
	case domain.FailureResponse:
		ctx.JSON(content.StatusCode, content.Response)
	case domain.SuccessResponse:
		ctx.JSON(content.StatusCode, content.Response)
	}
}

func ValidateEmail(ctx *gin.Context) {
	usecase := container.ContainerValidateEmail()
	data := usecase.Execute(ctx)
	switch content := data.(type) {
	case domain.FailureResponse:
		ctx.JSON(content.StatusCode, content.Response)
	case domain.SuccessResponse:
		ctx.JSON(content.StatusCode, content.Response)
	}
}

func ConfirmPhone(ctx *gin.Context) {
	var response interface{}
	usecaseValidateToken := container.ContainerValidateToken()
	usecaseConfirmPhone := container.ContainerConfirmPhone()
	data := usecaseValidateToken.Execute(ctx)
	switch content := data.(type) {
	case domain.FailureResponse:
		response = usecaseConfirmPhone.Execute(content.Response.Data, content.StatusCode)
	case domain.SuccessResponse:
		response = usecaseConfirmPhone.Execute(content.Response.Data, content.StatusCode)
	}

	switch content := response.(type) {
	case domain.FailureResponse:
		ctx.JSON(content.StatusCode, content.Response)
	case domain.SuccessResponse:
		ctx.JSON(content.StatusCode, content.Response)
	}
}
