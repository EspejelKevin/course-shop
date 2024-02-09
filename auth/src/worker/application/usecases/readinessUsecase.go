package usecases

import (
	"auth/src/worker/domain/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

type ReadinessUsecase struct {
	dbWorkerService repositories.DBRepository
}

func NewReadinessUsecase(dbWorkerService repositories.DBRepository) *ReadinessUsecase {
	return &ReadinessUsecase{
		dbWorkerService,
	}
}

func (readinessUsecase *ReadinessUsecase) Execute(ctx *gin.Context) map[string]interface{} {
	log.Println("Starting readiness usecase")
	status := readinessUsecase.dbWorkerService.IsUp()
	if status {
		log.Println("Mongo is up")
		return gin.H{
			"status":  200,
			"message": "Mongo is up",
		}
	}
	log.Println("Mongo is down")
	return gin.H{
		"status":  500,
		"message": "Mongo is down",
	}
}
