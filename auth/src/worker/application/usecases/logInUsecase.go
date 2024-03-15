package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/logger"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"auth/src/worker/domain/schemas"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var serviceLogInUsecase = "Login usecase"

type LogInUsecase struct {
	dbWorkerService repositories.DBRepository
	log             *logger.Log
}

func NewLogInUsecase(dbWorkerService repositories.DBRepository, log *logger.Log) *LogInUsecase {
	return &LogInUsecase{
		dbWorkerService,
		log,
	}
}

func (logInUsecase *LogInUsecase) Execute(ctx *gin.Context) interface{} {
	measurement := logger.Measurement{}
	measurement.Object = map[string]interface{}{}
	transactionId := uuid.NewString()
	logInUsecase.log.TracingId = transactionId
	logInUsecase.log.Info("Internal", serviceLogInUsecase, "Start login", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()
	_user, _ := ctx.Get("userIdentity")
	userIdentity := _user.(entities.UserIdentity)

	userDB := logInUsecase.dbWorkerService.GetUserByEmail(userIdentity.Email)

	if userDB == nil {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"email": userIdentity.Email}
		logInUsecase.log.Error("Internal", serviceLogInUsecase,
			"User does not exist", "Incorrect email", &measurement)
		data := map[string]interface{}{"user_message": "Invalid email"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
	}

	if !userDB.VerifiedEmail {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		measurement.Object = map[string]interface{}{"email": userIdentity.Email, "verified_email": userDB.VerifiedEmail}
		logInUsecase.log.Error("Internal", serviceLogInUsecase,
			"User not verified", "Email not verified", &measurement)
		data := map[string]interface{}{"user_message": "Please verify your email"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 403)
	}

	if !utils.CheckPasswordHash(userIdentity.Password, userDB.Password) {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		logInUsecase.log.Error("Internal", serviceLogInUsecase,
			"Passwords do not match", "Invalid password", &measurement)
		data := map[string]interface{}{"user_message": "Invalid password"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
	}

	userToMap := structs.Map(userDB)
	token, err := schemas.CreateAccessToken(userToMap)

	if err != nil {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement.TimeElapsed = timeElapsed
		logInUsecase.log.Error("Internal", serviceLogInUsecase,
			"Token not created", err.Error(), &measurement)
		data := map[string]interface{}{"user_message": "Failed to create access token"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement.TimeElapsed = timeElapsed
	logInUsecase.log.Info("Internal", serviceLogInUsecase, "User logged in successfully", &measurement)
	data := map[string]interface{}{"status": "User logged in successfully", "token": token}
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)

}
