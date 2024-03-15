package infrastructure

import (
	"auth/src/shared/domain"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func InternalServerErrorHandler(ctx *gin.Context, err any) {
	createHandler("Internal Server Error", 500, ctx)
}

func NotFoundErrorHandler(ctx *gin.Context) {
	createHandler("Resource Not Found", 404, ctx)
}

func MethodNotAllowedErrorHandler(ctx *gin.Context) {
	createHandler("Method Not Allowed", 405, ctx)
}

func createHandler(message string, statusCode int, ctx *gin.Context) {
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	data := map[string]interface{}{
		"user_message": message,
	}
	meta := domain.MetadataResponse{
		TransactionID: transactionId,
		Timestamp:     timestamp,
		TimeElapsed:   timeElapsed,
	}
	response := domain.Response{
		Data: data,
		Meta: meta,
	}
	failureResponse := domain.FailureResponse{
		StatusCode: statusCode,
		Response:   response,
	}

	ctx.AbortWithStatusJSON(failureResponse.StatusCode, failureResponse.Response)
}
