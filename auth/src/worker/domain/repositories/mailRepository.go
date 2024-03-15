package repositories

import (
	"auth/src/shared/logger"
	"auth/src/worker/domain/entities"
)

type MailRepository interface {
	IsUp(log *logger.Log) bool
	SendMail(email *entities.Email) error
}
