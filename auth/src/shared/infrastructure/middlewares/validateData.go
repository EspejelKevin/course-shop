package middlewares

import (
	"auth/src/shared/domain"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidatePayloadJSON(ctx *gin.Context) {
	log.Println("Starting middleware validatePayloadJSON")
	var userBody entities.User
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()

	if err := ctx.ShouldBindJSON(&userBody); err != nil {
		data := map[string]interface{}{
			"user_message": "Invalid parameters",
			"details":      utils.ValidationMsg(err),
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		response := domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
		content, _ := response.(domain.FailureResponse)
		ctx.JSON(content.StatusCode, content.Response)
		ctx.Abort()
		return
	}
	ctx.Next()
}
