package databases

import (
	"auth/src/shared/domain"
	"auth/src/worker/domain/entities"
	"log"
	"sync"

	"github.com/Masterminds/squirrel"
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

func (mysqlWorkerRepository *MySQLWorkerRepository) GetUserByEmail(email string) *entities.User {
	query, args, err := squirrel.Select("*").From("users").Where(squirrel.Eq{"email": email}).ToSql()
	var user entities.User
	if err != nil {
		log.Println("Failed to create sql query:", err)
		return nil
	}

	db := mysqlWorkerRepository.sessionFactory.GetDb()
	row := db.QueryRow(query, args...)
	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Lastname,
		&user.Password,
		&user.Email,
		&user.Verified,
		&user.Phone,
		&user.Rol,
	)
	if err != nil {
		log.Println("Failed to bind data to user:", err)
		return nil
	}

	return &user
}

func (mysqlWorkerRepository *MySQLWorkerRepository) CreateUser(user *entities.User) bool {
	query, args, err := squirrel.Insert("users").
		Columns("name", "lastname", "password", "email", "phone", "rol").
		Values(user.Name, user.Lastname, user.Password, user.Email, user.Phone, user.Rol).
		ToSql()
	if err != nil {
		log.Println("Failed to create sql query:", err)
		return false
	}

	db := mysqlWorkerRepository.sessionFactory.GetDb()
	_, err = db.Exec(query, args...)
	if err != nil {
		log.Println("Failed to insert user:", err)
		return false
	}

	return true
}
