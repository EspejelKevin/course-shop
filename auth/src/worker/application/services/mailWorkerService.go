package services

import (
	"auth/src/shared/logger"
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

func (mailWorkerService *MailWorkerService) IsUp(log *logger.Log) bool {
	return mailWorkerService.mail.IsUp(log)
}

func (mailWorkerService *MailWorkerService) SendMail(email *entities.Email) error {
	return mailWorkerService.mail.SendMail(email)
}
