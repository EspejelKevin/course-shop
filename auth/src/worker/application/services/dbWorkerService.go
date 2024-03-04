package services

import (
	"auth/src/worker/domain/entities"
	"auth/src/worker/domain/repositories"
)

type DBWorkerService struct {
	db repositories.DBRepository
}

func NewDBWorkerService(db repositories.DBRepository) *DBWorkerService {
	return &DBWorkerService{
		db,
	}
}

func (dbWorkerRepository *DBWorkerService) IsUp() bool {
	return dbWorkerRepository.db.IsUp()
}

func (dbWorkerRepository *DBWorkerService) CreateUser(user *entities.User) bool {
	return dbWorkerRepository.db.CreateUser(user)
}

func (dbWorkerRepository *DBWorkerService) GetUserByEmail(email string) *entities.User {
	return dbWorkerRepository.db.GetUserByEmail(email)
}

func (dbWorkerRepository *DBWorkerService) UpdateUserVerificationCode(email, code string) bool {
	return dbWorkerRepository.db.UpdateUserVerificationCode(email, code)
}
