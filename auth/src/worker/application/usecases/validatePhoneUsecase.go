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

var serviceValidatePhone = "Validate Phone usecase"

type ValidatePhoneUsecase struct {
	dbWorkerService repositories.DBRepository
	log             *logger.Log
}

func NewValidatePhoneUsecase(dbWorkerService repositories.DBRepository, log *logger.Log) *ValidatePhoneUsecase {
	return &ValidatePhoneUsecase{
		dbWorkerService,
		log,
	}
}

func (validatePhoneUsecase *ValidatePhoneUsecase) Execute(ctx *gin.Context) interface{} {
	measurement := logger.Measurement{}
	measurement.Object = map[string]interface{}{}
	transactionId := uuid.NewString()
	validatePhoneUsecase.log.TracingId = transactionId
	validatePhoneUsecase.log.Info("Internal", serviceValidatePhone, "Start validate phone", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()
	_code, _ := ctx.Get("code")
	code := _code.(string)
	codeEncoded := utils.Encode(code)

	result := validatePhoneUsecase.dbWorkerService.UpdateUserPhoneVerification(codeEncoded)

	if !result {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"codeEncoded": codeEncoded}
		validatePhoneUsecase.log.Error("Internal", serviceValidatePhone,
			"Error updating phone verification", "Phone already verified or code incorrect", &measurement)
		data := map[string]interface{}{"user_message": "Phone already verified"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 409)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement.TimeElapsed = timeElapsed
	validatePhoneUsecase.log.Info("Internal", serviceValidatePhone, "Phone verified", &measurement)
	data := map[string]interface{}{"status": "Phone verified successfully"}
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
