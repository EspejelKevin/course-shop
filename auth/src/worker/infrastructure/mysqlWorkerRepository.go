package infrastructure

import (
	"auth/src/shared/domain"
	"auth/src/worker/domain/entities"
	"log"
)

type MySQLWorkerRepository struct {
	db domain.Database
}

func NewMySQLWorkerRepository(db domain.Database) *MySQLWorkerRepository {
	return &MySQLWorkerRepository{
		db,
	}
}

func (mysqlWorkerRepository *MySQLWorkerRepository) IsUp() bool {
	data := mysqlWorkerRepository.db.IsUp()
	status := data["status"].(bool)
	message := data["message"].(string)

	if status {
		log.Println("Mongo is up", message)
	} else {
		log.Println("Mongo is down", message)
	}

	return status
}

func (mysqlWorkerRepository *MySQLWorkerRepository) GetUserByEmail(email string) map[string]interface{} {
	return map[string]interface{}{}
}

func (mysqlWorkerRepository *MySQLWorkerRepository) CreateUser(user *entities.User) bool {
	return true
}
