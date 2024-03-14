package repositories

import (
	"auth/src/shared/logger"
	"auth/src/worker/domain/entities"
)

type PhoneRepository interface {
	IsUp(log *logger.Log) bool
	SendMessage(message *entities.Message) error
}
