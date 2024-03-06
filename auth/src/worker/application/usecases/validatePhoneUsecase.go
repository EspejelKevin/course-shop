package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/utils"
	"auth/src/worker/domain/repositories"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ValidatePhoneUsecase struct {
	dbWorkerService repositories.DBRepository
}

func NewValidatePhoneUsecase(dbWorkerService repositories.DBRepository) *ValidatePhoneUsecase {
	return &ValidatePhoneUsecase{
		dbWorkerService,
	}
}

func (validatePhoneUsecase *ValidatePhoneUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting validate phone usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	_code, _ := ctx.Get("code")
	code := _code.(string)
	codeEncoded := utils.Encode(code)

	result := validatePhoneUsecase.dbWorkerService.UpdateUserPhoneVerification(codeEncoded)

	if !result {
		log.Println("Error verifying user. Check code: ", codeEncoded)
		data := map[string]interface{}{"user_message": "Phone already verified"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 409)
	}

	log.Println("Phone verified, code:", codeEncoded)
	data := map[string]interface{}{"status": "Phone verified successfully"}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
