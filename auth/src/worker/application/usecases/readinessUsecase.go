package usecases

import (
	"auth/src/worker/domain/repositories"
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

func (readinessUsecase *ReadinessUsecase) Execute(ctx *gin.Context) (map[string]interface{}, int) {
	log.Println("Starting readiness usecase")
	status := readinessUsecase.dbWorkerService.IsUp()
	transactionId := uuid.NewString()
	timestamp := time.Now().Format(time.Stamp)

	if !status {
		log.Println("Mongo is not up")
		return map[string]interface{}{
			"data": map[string]interface{}{
				"user_message": "Error de conexión a Mongo",
			},
			"meta": map[string]interface{}{
				"transaction_id": transactionId,
				"timestamp":      timestamp,
			},
		}, 500
	}

	return map[string]interface{}{
		"data": map[string]interface{}{
			"status": "Mongo is up",
		},
		"meta": map[string]interface{}{
			"transaction_id": transactionId,
			"timestamp":      timestamp,
		},
	}, 200
}
