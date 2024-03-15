package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/logger"
	"auth/src/shared/utils"
	"auth/src/worker/domain/repositories"
	"auth/src/worker/domain/schemas"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var serviceValidateToken = "Validate Token usecase"

type ValidateTokenUsecase struct {
	dbWorkerService repositories.DBRepository
	log             *logger.Log
}

func NewValidateTokenUsecase(dbWorkerService repositories.DBRepository, log *logger.Log) *ValidateTokenUsecase {
	return &ValidateTokenUsecase{
		dbWorkerService,
		log,
	}
}

func (validateTokenUsecase *ValidateTokenUsecase) Execute(ctx *gin.Context) interface{} {
	measurement := logger.Measurement{}
	measurement.Object = map[string]interface{}{}
	transactionId := uuid.NewString()
	validateTokenUsecase.log.TracingId = transactionId
	validateTokenUsecase.log.Info("Internal", serviceValidateToken, "Start validate token", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()
	_token, _ := ctx.Get("token")
	token := _token.(string)

	payload, err := schemas.ValidateAccessToken(token)
	payload = utils.Lower(payload).(map[string]interface{})

	if err != nil {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"token": token}
		validateTokenUsecase.log.Error("Internal", serviceValidateToken,
			"Token expired or incorrect", err.Error(), &measurement)
		data := map[string]interface{}{"user_message": "User unauthorized"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 401)
	}

	delete(payload, "password")
	delete(payload, "id")

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement.TimeElapsed = timeElapsed
	validateTokenUsecase.log.Info("Internal", serviceValidateToken, "Validated token successfully", &measurement)
	data := map[string]interface{}{"status": "Valid token", "payload": payload}
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
