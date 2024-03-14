package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/infrastructure"
	"auth/src/shared/logger"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

var serviceConfirmPhone = "Confirm Phone usecase"

type ConfirmPhoneUsecase struct {
	dbWorkerService    repositories.DBRepository
	phoneWorkerService repositories.PhoneRepository
	settings           *infrastructure.Settings
	log                *logger.Log
}

func NewConfirmPhoneUsecase(dbWorkerService repositories.DBRepository,
	phoneWorkerService repositories.PhoneRepository,
	settings *infrastructure.Settings, log *logger.Log) *ConfirmPhoneUsecase {
	return &ConfirmPhoneUsecase{
		dbWorkerService,
		phoneWorkerService,
		settings,
		log,
	}
}

func (confirmPhoneUsecase *ConfirmPhoneUsecase) Execute(response map[string]interface{}, statusCode int) interface{} {
	measurement := logger.Measurement{}
	measurement.Object = map[string]interface{}{}
	transactionId := uuid.NewString()
	confirmPhoneUsecase.log.TracingId = transactionId
	confirmPhoneUsecase.log.Info("Internal", serviceConfirmPhone, "Start confirm phone", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()

	if statusCode != 200 {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		confirmPhoneUsecase.log.Error("Internal", serviceConfirmPhone,
			"Error from validate token usecase", "Incorrect token or expired", &measurement)
		return domain.GenerateResponse(response, "failure", transactionId, timestamp, timeElapsed, statusCode)
	}

	phone := response["payload"].(map[string]interface{})["phone"].(string)
	name := response["payload"].(map[string]interface{})["name"].(string)
	email := response["payload"].(map[string]interface{})["email"].(string)
	code := randstr.String(20)
	codeEncoded := utils.Encode(code)

	result := confirmPhoneUsecase.dbWorkerService.UpdateUserPhoneVerificationCode(email, codeEncoded)

	if !result {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"codeEncoded": codeEncoded, "email": email}
		confirmPhoneUsecase.log.Error("Internal", serviceConfirmPhone,
			"Error to update user phone verification",
			"Incorrect code or phone already verified", &measurement)
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
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"code": code, "email": email}
		confirmPhoneUsecase.log.Error("Internal", serviceConfirmPhone,
			"Error sending message",
			err.Error(), &measurement)
		data := map[string]interface{}{
			"internal_message": "Error sending message",
		}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement.TimeElapsed = timeElapsed
	confirmPhoneUsecase.log.Info("Internal", serviceConfirmPhone, "Message sent", &measurement)
	data := map[string]interface{}{
		"status":  "Message sent successfully",
		"message": "We sent an message with your verification code",
	}
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
