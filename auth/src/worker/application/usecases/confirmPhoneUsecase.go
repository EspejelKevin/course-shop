package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/infrastructure"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type ConfirmPhoneUsecase struct {
	dbWorkerService    repositories.DBRepository
	phoneWorkerService repositories.PhoneRepository
	settings           *infrastructure.Settings
}

func NewConfirmPhoneUsecase(dbWorkerService repositories.DBRepository,
	phoneWorkerService repositories.PhoneRepository,
	settings *infrastructure.Settings) *ConfirmPhoneUsecase {
	return &ConfirmPhoneUsecase{
		dbWorkerService,
		phoneWorkerService,
		settings,
	}
}

func (confirmPhoneUsecase *ConfirmPhoneUsecase) Execute(response map[string]interface{}, statusCode int) interface{} {
	log.Println("Starting confirm phone usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()

	if statusCode != 200 {
		log.Println("Error from validate token usecase")
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(response, "failure", transactionId, timestamp, timeElapsed, statusCode)
	}

	phone := response["payload"].(map[string]interface{})["phone"].(string)
	name := response["payload"].(map[string]interface{})["name"].(string)
	email := response["payload"].(map[string]interface{})["email"].(string)
	code := randstr.String(20)
	codeEncoded := utils.Encode(code)

	result := confirmPhoneUsecase.dbWorkerService.UpdateUserPhoneVerificationCode(email, codeEncoded)

	if !result {
		log.Println("Error to update user phone verification code")
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		data := map[string]interface{}{"user_message": "Code verification operation failed. Please try again"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	message := entities.Message{
		To:   fmt.Sprintf("whatsapp:%s", phone),
		From: confirmPhoneUsecase.settings.PhoneFrom,
		Body: fmt.Sprintf("Hi %s, this is your verification code: *%s*. For your security, do not share your code.", name, code),
	}

	err := confirmPhoneUsecase.phoneWorkerService.SendMessage(&message)

	if err != nil {
		log.Println("Error sending message:", err)
		data := map[string]interface{}{
			"internal_message": "Error sending message",
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	log.Println("Message sent successfully")
	data := map[string]interface{}{
		"status":  "Message sent successfully",
		"message": "We sent an message with your verification code",
	}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
