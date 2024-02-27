package usecases

import (
	"auth/src/worker/domain/repositories"

	"github.com/gin-gonic/gin"
)

type SignInUsecase struct {
	dbWorkerService repositories.DBRepository
}

func NewSignUpUsecase(dbWorkerService repositories.DBRepository) *SignInUsecase {
	return &SignInUsecase{
		dbWorkerService,
	}
}

func (signUpUsecase *SignInUsecase) Execute(ctx *gin.Context) map[string]interface{} {
	return map[string]interface{}{}
}
