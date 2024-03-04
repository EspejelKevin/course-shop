package services

import (
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
)

type MailWorkerService struct {
	mail repositories.MailRepository
}

func NewMailWorkerService(mail repositories.MailRepository) *MailWorkerService {
	return &MailWorkerService{
		mail,
	}
}

func (mailWorkerService *MailWorkerService) IsUp() bool {
	return mailWorkerService.mail.IsUp()
}

func (mailWorkerService *MailWorkerService) SendMail(email *entities.Email) error {
	return mailWorkerService.mail.SendMail(email)
}
