package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/infrastructure"
	"auth/src/shared/logger"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var serviceConfirmEmail = "Confirm Email usecase"

type ConfirmEmailUsecase struct {
	settings          *infrastructure.Settings
	mailWorkerService repositories.MailRepository
	log               *logger.Log
}

func NewConfirmEmailUsecase(mailWorkerService repositories.MailRepository,
	settings *infrastructure.Settings, log *logger.Log) *ConfirmEmailUsecase {
	return &ConfirmEmailUsecase{
		settings,
		mailWorkerService,
		log,
	}
}

func (confirmEmailUsecase *ConfirmEmailUsecase) Execute(ctx *gin.Context) interface{} {
	measurement := logger.Measurement{}
	measurement.Object = map[string]interface{}{}
	transactionId := uuid.NewString()
	confirmEmailUsecase.log.TracingId = transactionId
	confirmEmailUsecase.log.Info("Internal", serviceConfirmEmail, "Start confirm email", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()
	_emailData, _ := ctx.Get("emailData")
	emailData := _emailData.(entities.Email)

	err := confirmEmailUsecase.mailWorkerService.SendMail(&emailData)

	if err != nil {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"code": emailData.Code, "emailTo": emailData.To}
		confirmEmailUsecase.log.Error("Internal", serviceConfirmEmail,
			"Error sending email", err.Error(), &measurement)
		data := map[string]interface{}{
			"internal_message": "Error sending email",
		}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement.TimeElapsed = timeElapsed
	confirmEmailUsecase.log.Info("Internal", serviceConfirmEmail, "Email sent", &measurement)
	data := map[string]interface{}{
		"status":  "Email sent successfully",
		"message": "We sent an email with your verification code",
	}
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)

}
