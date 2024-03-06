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

type ValidateEmailUsecase struct {
	dbWorkerService repositories.DBRepository
}

func NewValidateEmailUsecase(dbWorkerService repositories.DBRepository) *ValidateEmailUsecase {
	return &ValidateEmailUsecase{
		dbWorkerService,
	}
}

func (validateEmailUsecase *ValidateEmailUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting validate email usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	_code, _ := ctx.Get("code")
	code := _code.(string)
	codeEncoded := utils.Encode(code)

	result := validateEmailUsecase.dbWorkerService.UpdateUserEmailVerification(codeEncoded)

	if !result {
		log.Println("Error verifying user. Check code: ", codeEncoded)
		data := map[string]interface{}{"user_message": "Email already verified"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 409)
	}

	log.Println("User verified, code:", codeEncoded)
	data := map[string]interface{}{"status": "Email verified successfully"}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
