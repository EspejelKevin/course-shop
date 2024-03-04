package usecases

import (
	"auth/src/shared/domain"
	"auth/src/shared/infrastructure"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type SignUpUsecase struct {
	dbWorkerService   repositories.DBRepository
	settings          *infrastructure.Settings
	mailWorkerService repositories.MailRepository
}

func NewSignUpUsecase(dbWorkerService repositories.DBRepository, mailWorkerService repositories.MailRepository,
	settings *infrastructure.Settings) *SignUpUsecase {
	return &SignUpUsecase{
		dbWorkerService,
		settings,
		mailWorkerService,
	}
}

func (signUpUsecase *SignUpUsecase) Execute(ctx *gin.Context) interface{} {
	log.Println("Starting signUp usecase")
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
		data := map[string]interface{}{"user_message": "User not created. Please try again"}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	code := randstr.String(20)
	codeEncoded := utils.Encode(code)
	result = signUpUsecase.dbWorkerService.UpdateUserVerificationCode(user.Email, codeEncoded)

	if !result {
		log.Println("Error to update user verification code")
		data := map[string]interface{}{
			"internal_message": "Error to update user verification code",
			"code":             code,
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	emailData := entities.Email{
		URL:     signUpUsecase.settings.URLValidateMail,
		Name:    user.Name,
		Subject: "Verification Code",
		From:    signUpUsecase.settings.EmailFrom,
		To:      user.Email,
		Code:    code,
	}

	err := signUpUsecase.mailWorkerService.SendMail(&emailData)

	if err != nil {
		log.Println("Error sending email:", err)
		data := map[string]interface{}{
			"internal_message": "Error sending email",
			"code":             code,
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		return domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 500)
	}

	log.Println("User created successfully")
	data := map[string]interface{}{
		"status":  "User created successfully",
		"message": "We sent an email with your verification code",
	}
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	return domain.GenerateResponse(data, "", transactionId, timestamp, timeElapsed, 200)
}
