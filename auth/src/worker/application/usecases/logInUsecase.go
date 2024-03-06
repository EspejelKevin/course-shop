package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"auth/src/worker/domain/schemas"
	"fmt"
	"log"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LogInUsecase struct {
	dbWorkerService repositories.DBRepository
}

func NewLogInUsecase(dbWorkerService repositories.DBRepository) *LogInUsecase {
	return &LogInUsecase{
		dbWorkerService,
	}
}

func (logInUsecase *LogInUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting logIn usecase")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	_user, _ := ctx.Get("userIdentity")
	userIdentity := _user.(entities.UserIdentity)

	userDB := logInUsecase.dbWorkerService.GetUserByEmail(userIdentity.Email)

	if userDB == nil {
		log.Println("User does not exist")
		data := map[string]interface{}{"user_message": "Invalid email"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
	}

	if !userDB.VerifiedEmail {
		log.Println("User not verified")
		data := map[string]interface{}{"user_message": "Please verify your email"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 403)
	}

	if !utils.CheckPasswordHash(userIdentity.Password, userDB.Password) {
		log.Println("The hashed password does not match the password from request")
		data := map[string]interface{}{"user_message": "Invalid password"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
	}

	userToMap := structs.Map(userDB)
	token, err := schemas.CreateAccessToken(userToMap)

	if err != nil {
		log.Println("Could not create access token:", err)
		data := map[string]interface{}{"user_message": "Failed to create access token"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	log.Println("User logged in successfully")
	data := map[string]interface{}{"status": "User logged in successfully", "token": token}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)

}
