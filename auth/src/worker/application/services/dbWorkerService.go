package services

import (
	"auth/src/shared/logger"
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

func (dbWorkerService *DBWorkerService) IsUp(log *logger.Log) bool {
	return dbWorkerService.db.IsUp(log)
}

func (dbWorkerService *DBWorkerService) CreateUser(user *entities.User) bool {
	return dbWorkerService.db.CreateUser(user)
}

func (dbWorkerService *DBWorkerService) GetUserByEmail(email string) *entities.User {
	return dbWorkerService.db.GetUserByEmail(email)
}

func (dbWorkerService *DBWorkerService) UpdateUserEmailVerificationCode(email, code string) bool {
	return dbWorkerService.db.UpdateUserEmailVerificationCode(email, code)
}

func (dbWorkerService *DBWorkerService) UpdateUserEmailVerification(code string) bool {
	return dbWorkerService.db.UpdateUserEmailVerification(code)
}

func (dbWorkerService *DBWorkerService) UpdateUserPhoneVerificationCode(email, code string) bool {
	return dbWorkerService.db.UpdateUserPhoneVerificationCode(email, code)
}

func (dbWorkerService *DBWorkerService) UpdateUserPhoneVerification(code string) bool {
	return dbWorkerService.db.UpdateUserPhoneVerification(code)
}
