package databases

import (
	"auth/src/shared/domain"
	"auth/src/worker/domain/entities"
	"log"
	"sync"
)

var once sync.Once
var mysqlWorkerRepository *MySQLWorkerRepository

type MySQLWorkerRepository struct {
	sessionFactory domain.Database
}

func NewMySQLWorkerRepository(sessionFactory domain.Database) *MySQLWorkerRepository {
	once.Do(func() {
		mysqlWorkerRepository = &MySQLWorkerRepository{
			sessionFactory,
		}
	})
	return mysqlWorkerRepository
}

func (mysqlWorkerRepository *MySQLWorkerRepository) IsUp() bool {
	data := mysqlWorkerRepository.sessionFactory.IsUp()
	status := data["status"].(bool)
	message := data["message"].(string)

	if status {
		log.Println("MySQL is up", message)
	} else {
		log.Println("MySQL is down", message)
	}

	return status
}

func (mysqlWorkerRepository *MySQLWorkerRepository) GetUserByEmail(email string) map[string]interface{} {
	return map[string]interface{}{}
}

func (mysqlWorkerRepository *MySQLWorkerRepository) CreateUser(user *entities.User) bool {
	return true
}
