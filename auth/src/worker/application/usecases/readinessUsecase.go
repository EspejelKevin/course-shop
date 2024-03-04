package usecases

import (
	"auth/src/shared/domain"
	"auth/src/worker/domain/repositories"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReadinessUsecase struct {
	dbWorkerService   repositories.DBRepository
	mailWorkerService repositories.MailRepository
}

func NewReadinessUsecase(dbWorkerService repositories.DBRepository, mailWorkerService repositories.MailRepository) *ReadinessUsecase {
	return &ReadinessUsecase{
		dbWorkerService,
		mailWorkerService,
	}
}

func (readinessUsecase *ReadinessUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting readiness usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	status := readinessUsecase.dbWorkerService.IsUp()
	data := map[string]interface{}{"status": "MySQL and SMTP are up"}

	if !status {
		log.Println("MySQL is not up")
		data = map[string]interface{}{"user_message": "MySQL is not up"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	status = readinessUsecase.mailWorkerService.IsUp()

	if !status {
		log.Println("SMTP is not up")
		data = map[string]interface{}{"user_message": "SMTP is not up"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
