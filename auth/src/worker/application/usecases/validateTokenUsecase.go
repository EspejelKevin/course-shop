package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/utils"
	"auth/src/worker/domain/repositories"
	"auth/src/worker/domain/schemas"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ValidateTokenUsecase struct {
	dbWorkerService repositories.DBRepository
}

func NewValidateTokenUsecase(dbWorkerService repositories.DBRepository) *ValidateTokenUsecase {
	return &ValidateTokenUsecase{
		dbWorkerService,
	}
}

func (validateTokenUsecase *ValidateTokenUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting validate token usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	_token, _ := ctx.Get("token")
	token := _token.(string)

	payload, err := schemas.ValidateAccessToken(token)
	payload = utils.Lower(payload).(map[string]interface{})

	if err != nil {
		log.Println("Token error:", err)
		data := map[string]interface{}{"user_message": "User unauthorized"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 401)
	}

	delete(payload, "password")

	log.Println("Validated token successfully")
	data := map[string]interface{}{"status": "Valid token", "payload": payload}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
