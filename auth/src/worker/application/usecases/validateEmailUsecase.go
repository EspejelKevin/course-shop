package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/logger"
	"auth/src/shared/utils"
	"auth/src/worker/domain/repositories"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var serviceValidateEmail = "Validate Email usecase"

type ValidateEmailUsecase struct {
	dbWorkerService repositories.DBRepository
	log             *logger.Log
}

func NewValidateEmailUsecase(dbWorkerService repositories.DBRepository, log *logger.Log) *ValidateEmailUsecase {
	return &ValidateEmailUsecase{
		dbWorkerService,
		log,
	}
}

func (validateEmailUsecase *ValidateEmailUsecase) Execute(ctx *gin.Context) interface{} {
	measurement := logger.Measurement{}
	measurement.Object = map[string]interface{}{}
	transactionId := uuid.NewString()
	validateEmailUsecase.log.TracingId = transactionId
	validateEmailUsecase.log.Info("Internal", serviceValidateEmail, "Start validate email", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()
	_code, _ := ctx.Get("code")
	code := _code.(string)
	codeEncoded := utils.Encode(code)

	result := validateEmailUsecase.dbWorkerService.UpdateUserEmailVerification(codeEncoded)

	if !result {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"codeEncoded": codeEncoded}
		validateEmailUsecase.log.Error("Internal", serviceValidateEmail,
			"Error updating email verification", "Email already verified or code incorrect", &measurement)
		data := map[string]interface{}{"user_message": "Email already verified"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 409)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement.TimeElapsed = timeElapsed
	validateEmailUsecase.log.Info("Internal", serviceValidateEmail, "Email verified", &measurement)
	data := map[string]interface{}{"status": "Email verified successfully"}
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
