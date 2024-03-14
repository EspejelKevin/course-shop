package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/logger"
	"auth/src/worker/domain/repositories"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var serviceReadiness = "Readiness usecase"

type ReadinessUsecase struct {
	dbWorkerService    repositories.DBRepository
	mailWorkerService  repositories.MailRepository
	phoneWorkerService repositories.PhoneRepository
	log                *logger.Log
}

func NewReadinessUsecase(dbWorkerService repositories.DBRepository,
	mailWorkerService repositories.MailRepository,
	phoneWorkerService repositories.PhoneRepository,
	log *logger.Log) *ReadinessUsecase {
	return &ReadinessUsecase{
		dbWorkerService,
		mailWorkerService,
		phoneWorkerService,
		log,
	}
}

func (readinessUsecase *ReadinessUsecase) Execute(ctx *gin.Context) interface{} {
	measurement := logger.Measurement{}
	measurement.Object = map[string]interface{}{}
	transactionId := uuid.NewString()
	readinessUsecase.log.TracingId = transactionId
	readinessUsecase.log.Info("Internal", serviceReadiness, "Start readiness", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()
	status := readinessUsecase.dbWorkerService.IsUp(readinessUsecase.log)
	data := map[string]interface{}{"status": "MySQL, SMTP and PhoneServer are up"}

	if !status {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		readinessUsecase.log.Error("Internal",
			serviceReadiness, "MySQL is not up",
			"Erro to connect to MySQL", &measurement)
		data = map[string]interface{}{"user_message": "MySQL is not up"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	status = readinessUsecase.mailWorkerService.IsUp(readinessUsecase.log)

	if !status {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		readinessUsecase.log.Error("Internal",
			serviceReadiness, "SMTP is not up",
			"Erro to connect to SMTP Gmail. Send email error.", &measurement)
		data = map[string]interface{}{"user_message": "SMTP is not up"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	status = readinessUsecase.phoneWorkerService.IsUp(readinessUsecase.log)

	if !status {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		readinessUsecase.log.Error("Internal",
			serviceReadiness, "Phone server is not up",
			"Erro to connect to Twilio. Send whatsapp message.", &measurement)
		data = map[string]interface{}{"user_message": "Phone Server is not up"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement.TimeElapsed = timeElapsed
	readinessUsecase.log.Info("Internal", serviceReadiness, "MySQL, SMTP, Twilio are up", &measurement)
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
