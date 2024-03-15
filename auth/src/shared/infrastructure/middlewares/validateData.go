package middlewares

import (
	"auth/src/shared/domain"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var invalidParameters = "Invalid parameters"

func ValidatePayloadSignIn(ctx *gin.Context) {
	log.Println("Starting middleware ValidatePayloadSignUp")
	var userBody entities.User
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()

	if err := ctx.ShouldBindJSON(&userBody); err != nil {
		data := map[string]interface{}{
			"user_message": invalidParameters,
			"details":      utils.ValidationMsg(err),
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		response := domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
		content, _ := response.(domain.FailureResponse)
		ctx.JSON(content.StatusCode, content.Response)
		ctx.Abort()
		return
	}
	ctx.Set("user", userBody)
	ctx.Next()
}

func ValidatePayloadLogIn(ctx *gin.Context) {
	log.Println("Starting middleware ValidatePayloadLogIn")
	var userIdentity entities.UserIdentity
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()

	if err := ctx.ShouldBindJSON(&userIdentity); err != nil {
		data := map[string]interface{}{
			"user_message": invalidParameters,
			"details":      utils.ValidationMsg(err),
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		response := domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
		content, _ := response.(domain.FailureResponse)
		ctx.JSON(content.StatusCode, content.Response)
		ctx.Abort()
		return
	}
	ctx.Set("userIdentity", userIdentity)
	ctx.Next()
}

func ValidateBearerToken(ctx *gin.Context) {
	log.Println("Starting middleware ValidateBearerToken")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()

	authorizationHeader := ctx.GetHeader("Authorization")
	token := strings.Split(authorizationHeader, " ")

	if len(token) < 2 {
		data := map[string]interface{}{
			"user_message": "User unauthorized",
			"details":      []string{"missing header 'Authorization' or bad format bearer token"},
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		response := domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 401)
		content, _ := response.(domain.FailureResponse)
		ctx.JSON(content.StatusCode, content.Response)
		ctx.Abort()
		return
	}

	ctx.Set("token", token[1])
	ctx.Next()
}

func ValidateVerificationCode(ctx *gin.Context) {
	log.Println("Starting middleware ValidateVerificationCode")
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()
	var verificationCode entities.VerificationCode

	if err := ctx.ShouldBindJSON(&verificationCode); err != nil {
		data := map[string]interface{}{
			"user_message": invalidParameters,
			"details":      utils.ValidationMsg(err),
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		response := domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
		content, _ := response.(domain.FailureResponse)
		ctx.JSON(content.StatusCode, content.Response)
		ctx.Abort()
		return
	}
	ctx.Set("code", verificationCode.Code)
	ctx.Next()
}

func ValidateEmailData(ctx *gin.Context) {
	log.Println("Starting middleware ValidateEmailData")
	var emailData entities.Email
	timestamp := time.Now().Format(time.Stamp)
	transactionId := uuid.NewString()
	start := time.Now()

	if err := ctx.ShouldBindJSON(&emailData); err != nil {
		data := map[string]interface{}{
			"user_message": invalidParameters,
			"details":      utils.ValidationMsg(err),
		}
		timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
		response := domain.GenerateResponse(data, "failure", transactionId, timestamp, timeElapsed, 400)
		content, _ := response.(domain.FailureResponse)
		ctx.JSON(content.StatusCode, content.Response)
		ctx.Abort()
		return
	}
	ctx.Set("emailData", emailData)
	ctx.Next()
}
