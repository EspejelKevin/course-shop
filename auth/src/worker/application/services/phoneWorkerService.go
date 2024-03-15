package services

import (
	"auth/src/shared/logger"
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
)

type PhoneWorkerService struct {
	phoneServer repositories.PhoneRepository
}

func NewPhoneWorkerService(phoneServer repositories.PhoneRepository) *PhoneWorkerService {
	return &PhoneWorkerService{
		phoneServer,
	}
}

func (phoneWorkerService *PhoneWorkerService) IsUp(log *logger.Log) bool {
	return phoneWorkerService.phoneServer.IsUp(log)
}

func (phoneWorkerService *PhoneWorkerService) SendMessage(message *entities.Message) error {
	return phoneWorkerService.phoneServer.SendMessage(message)
}
