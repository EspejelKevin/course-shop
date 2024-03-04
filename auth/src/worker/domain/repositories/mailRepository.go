package repositories

import "auth/src/worker/domain/entities"

type MailRepository interface {
	IsUp() bool
	SendMail(email *entities.Email) error
}
