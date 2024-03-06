package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/infrastructure"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConfirmEmailUsecase struct {
	settings          *infrastructure.Settings
	mailWorkerService repositories.MailRepository
}

func NewConfirmEmailUsecase(mailWorkerService repositories.MailRepository,
	settings *infrastructure.Settings) *ConfirmEmailUsecase {
	return &ConfirmEmailUsecase{
		settings,
		mailWorkerService,
	}
}

func (confirmEmailUsecase *ConfirmEmailUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting confirm email usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	_emailData, _ := ctx.Get("emailData")
	emailData := _emailData.(entities.Email)

	err := confirmEmailUsecase.mailWorkerService.SendMail(&emailData)

	if err != nil {
		log.Println("Error sending email:", err)
		data := map[string]interface{}{
			"internal_message": "Error sending email",
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	log.Println("Email sent successfully")
	data := map[string]interface{}{
		"status":  "Email sent successfully",
		"message": "We sent an email with your verification code",
	}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)

}
