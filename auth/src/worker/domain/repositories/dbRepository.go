package repositories

import "auth/src/worker/domain/entities"

type DBRepository interface {
	IsUp() bool
	GetUserByEmail(email string) *entities.User
	UpdateUserVerification(code string) bool
	CreateUser(user *entities.User) bool
	UpdateUserVerificationCode(email, code string) bool
}
