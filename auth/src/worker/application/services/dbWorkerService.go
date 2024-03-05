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

func (dbWorkerService *DBWorkerService) IsUp() bool {
	return dbWorkerService.db.IsUp()
}

func (dbWorkerService *DBWorkerService) CreateUser(user *entities.User) bool {
	return dbWorkerService.db.CreateUser(user)
}

func (dbWorkerService *DBWorkerService) GetUserByEmail(email string) *entities.User {
	return dbWorkerService.db.GetUserByEmail(email)
}

func (dbWorkerService *DBWorkerService) UpdateUserVerificationCode(email, code string) bool {
	return dbWorkerService.db.UpdateUserVerificationCode(email, code)
}

func (dbWorkerService *DBWorkerService) UpdateUserVerification(code string) bool {
	return dbWorkerService.db.UpdateUserVerification(code)
}
