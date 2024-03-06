package repositories

import "auth/src/worker/domain/entities"

type PhoneRepository interface {
	IsUp() bool
	SendMessage(message *entities.Message) error
}
