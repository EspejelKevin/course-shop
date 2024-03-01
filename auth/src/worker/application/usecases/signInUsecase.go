package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignInUsecase struct {
	dbWorkerService repositories.DBRepository
}

func NewSignUpUsecase(dbWorkerService repositories.DBRepository) *SignInUsecase {
	return &SignInUsecase{
		dbWorkerService,
	}
}

func (signUpUsecase *SignInUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting signIn usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	_user, _ := ctx.Get("user")
	user := _user.(entities.User)

	userDB := signUpUsecase.dbWorkerService.GetUserByEmail(user.Email)

	if userDB != nil {
		log.Println("User already exists")
		data := map[string]interface{}{"user_message": "User already exists"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 409)
	}

	user.Password = utils.HashPassword(user.Password)
	result := signUpUsecase.dbWorkerService.CreateUser(&user)

	if !result {
		log.Println("Error to create user")
		data := map[string]interface{}{"user_message": "User not created"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	log.Println("User created successfully")
	data := map[string]interface{}{"status": "User created successfully"}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
