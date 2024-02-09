package services

import (
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
)

type DBWorkerService struct {
	sessionFactory repositories.DBRepository
}

func NewDBWorkerService(sessionFactory repositories.DBRepository) *DBWorkerService {
	return &DBWorkerService{
		sessionFactory,
	}
}

func (dbWorkerRepository *DBWorkerService) IsUp() bool {
	return dbWorkerRepository.sessionFactory.IsUp()
}

func (dbWorkerRepository *DBWorkerService) CreateUser(user *entities.User) bool {
	return dbWorkerRepository.sessionFactory.CreateUser(user)
}

func (dbWorkerRepository *DBWorkerService) GetUserByEmail(email string) map[string]interface{} {
	return dbWorkerRepository.sessionFactory.GetUserByEmail(email)
}
