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
	dbWorkerService repositories.DBRepository
}

func NewReadinessUsecase(dbWorkerService repositories.DBRepository) *ReadinessUsecase {
	return &ReadinessUsecase{
		dbWorkerService,
	}
}

func (readinessUsecase *ReadinessUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting readiness usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	status := readinessUsecase.dbWorkerService.IsUp()
	data := map[string]interface{}{"status": "MySQL is up"}

	if !status {
		log.Println("MySQL is not up")
		data = map[string]interface{}{"user_message": "MySQL is not up"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
