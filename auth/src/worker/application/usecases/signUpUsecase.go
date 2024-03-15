package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/infrastructure"
	"auth/src/shared/logger"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

var serviceSignUp = "Sign up usecase"

type SignUpUsecase struct {
	dbWorkerService   repositories.DBRepository
	settings          *infrastructure.Settings
	mailWorkerService repositories.MailRepository
	log               *logger.Log
}

func NewSignUpUsecase(dbWorkerService repositories.DBRepository, mailWorkerService repositories.MailRepository,
	settings *infrastructure.Settings, log *logger.Log) *SignUpUsecase {
	return &SignUpUsecase{
		dbWorkerService,
		settings,
		mailWorkerService,
		log,
	}
}

func (signUpUsecase *SignUpUsecase) Execute(ctx *gin.Context) interface{} {
	transactionId := uuid.NewString()
	signUpUsecase.log.TracingId = transactionId
	signUpUsecase.log.Info("Internal", serviceSignUp, "Start sign up", nil)
	timestamp := time.Now().Format(time.Stamp)
	start := time.Now()
	_user, _ := ctx.Get("user")
	user := _user.(entities.User)

	userDB := signUpUsecase.dbWorkerService.GetUserByEmail(user.Email)

	if userDB != nil {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement := logger.Measurement{
			TimeElapsed: timeElapsed,
			Object:      map[string]interface{}{"user": user.Email, "id": user.Id, "rol": user.Rol},
		}
		signUpUsecase.log.Error("Internal", serviceSignUp, "User already exists", "Duplicate user", &measurement)
		data := map[string]interface{}{"user_message": "User already exists"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 409)
	}

	user.Password = utils.HashPassword(user.Password)
	result := signUpUsecase.dbWorkerService.CreateUser(&user)

	if !result {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement := logger.Measurement{
			TimeElapsed: timeElapsed,
			Object:      map[string]interface{}{"user": user.Email, "id": user.Id, "rol": user.Rol},
		}
		signUpUsecase.log.Error("Internal", serviceSignUp,
			"Error in MySQL. Create user",
			"Query failed.", &measurement)
		data := map[string]interface{}{"user_message": "User not created. Please try again"}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	code := randstr.String(20)
	codeEncoded := utils.Encode(code)
	result = signUpUsecase.dbWorkerService.UpdateUserEmailVerificationCode(user.Email, codeEncoded)

	if !result {
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement := logger.Measurement{
			TimeElapsed: timeElapsed,
			Object:      map[string]interface{}{"user": user.Email, "id": user.Id, "code": codeEncoded},
		}
		signUpUsecase.log.Error("Internal", serviceSignUp,
			"Error in MySQL. Update user email verification code",
			"Query failed to update", &measurement)
		data := map[string]interface{}{
			"internal_message": "Error to update user verification code",
			"code":             code,
		}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	emailData := entities.Email{
		Name:    user.Name,
		Subject: "Verification Code",
		From:    signUpUsecase.settings.EmailFrom,
		To:      user.Email,
		Code:    code,
	}

	err := signUpUsecase.mailWorkerService.SendMail(&emailData)

	if err != nil {
		payload := structs.Map(emailData)
		payload = utils.Lower(payload).(map[string]interface{})
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		measurement := logger.Measurement{
			TimeElapsed: timeElapsed,
			Object: map[string]interface{}{
				"user": user.Email,
				"code": code,
			},
		}
		signUpUsecase.log.Error("Internal", serviceSignUp,
			"Error in SMTP Gmail. Send email",
			err.Error(), &measurement)
		data := map[string]interface{}{
			"internal_message": "Error sending email",
			"email_data":       payload,
		}
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement := logger.Measurement{
		TimeElapsed: timeElapsed,
		Object:      map[string]interface{}{"user": user.Email},
	}
	signUpUsecase.log.Info("Internal", serviceSignUp, "User created", &measurement)
	data := map[string]interface{}{
		"status":  "User created successfully",
		"message": "We sent an email with your verification code",
	}
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
