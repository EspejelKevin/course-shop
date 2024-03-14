package repositories

import (
	"auth/src/shared/logger"
	"auth/src/worker/domain/entities"
)

type DBRepository interface {
	IsUp(log *logger.Log) bool
	GetUserByEmail(email string) *entities.User
	CreateUser(user *entities.User) bool
	UpdateUserEmailVerification(code string) bool
	UpdateUserEmailVerificationCode(email, code string) bool
	UpdateUserPhoneVerification(code string) bool
	UpdateUserPhoneVerificationCode(email, code string) bool
}
