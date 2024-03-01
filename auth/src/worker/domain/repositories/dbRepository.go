package repositories

import "auth/src/worker/domain/entities"

type DBRepository interface {
	IsUp() bool
	GetUserByEmail(email string) *entities.User
	CreateUser(user *entities.User) bool
}
