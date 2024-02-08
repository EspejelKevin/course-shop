package repositories

import "auth/src/worker/domain/entities"

type DBRepository interface {
	IsUp() bool
	GetUserByEmail(email string) map[string]interface{}
	CreateUser(user *entities.User) bool
}
