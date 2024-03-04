package usecases

import (
	"auth/src/worker/domain/repositories"
	"log"

	"github.com/gin-gonic/gin"
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
	// timestamp := time.Now().Format(time.Stamp)
	// transactionId := uuid.NewString()
	// start := time.Now()
	// _email, _ := ctx.Get("email")
	// email := _email.(string)

	return "Cuenta verificada"
}
