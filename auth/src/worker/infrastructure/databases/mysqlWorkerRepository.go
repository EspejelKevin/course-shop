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
var messageError = "Failed to create sql query:"
var messageUpdate = "Failed to update user:"
var messageBindData = "Failed to bind data to user:"

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
	query, args, err := squirrel.Select("id", "name", "lastname", "password", "email", "verifiedemail", "phone", "rol").
		From("users").
		Where(squirrel.Eq{"email": email}).
		ToSql()
	var user entities.User
	if err != nil {
		log.Println(messageError, err)
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
		&user.VerifiedEmail,
		&user.Phone,
		&user.Rol,
	)
	if err != nil {
		log.Println(messageBindData, err)
		return nil
	}

	return &user
}

func (mysqlWorkerRepository *MySQLWorkerRepository) UpdateUserEmailVerification(code string) bool {
	query, args, err := squirrel.Select("verifiedemail", "email").
		From("users").
		Where(squirrel.Eq{"codeemail": code}).
		ToSql()
	var verified bool
	var email string
	if err != nil {
		log.Println(messageError, err)
		return false
	}

	db := mysqlWorkerRepository.sessionFactory.GetDb()
	row := db.QueryRow(query, args...)
	err = row.Scan(&verified, &email)
	if err != nil {
		log.Println(messageBindData, err)
		return false
	}

	if verified {
		log.Println("User verified with mail, ", email)
		return false
	}

	query, args, err = squirrel.Update("users").
		Set("codeemail", "").
		Set("verifiedemail", true).
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		log.Println(messageError, err)
		return false
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		log.Println(messageUpdate, err)
		return false
	}

	return true
}

func (mysqlWorkerRepository *MySQLWorkerRepository) CreateUser(user *entities.User) bool {
	query, args, err := squirrel.Insert("users").
		Columns("name", "lastname", "password", "email", "phone", "rol").
		Values(user.Name, user.Lastname, user.Password, user.Email, user.Phone, user.Rol).
		ToSql()
	if err != nil {
		log.Println(messageError, err)
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

func (mysqlWorkerRepository *MySQLWorkerRepository) UpdateUserEmailVerificationCode(email, code string) bool {
	query, args, err := squirrel.Update("users").
		Set("codeemail", code).
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		log.Println(messageError, err)
		return false
	}

	db := mysqlWorkerRepository.sessionFactory.GetDb()
	_, err = db.Exec(query, args...)
	if err != nil {
		log.Println(messageUpdate, err)
		return false
	}

	return true
}

func (mysqlWorkerRepository *MySQLWorkerRepository) UpdateUserPhoneVerificationCode(email, code string) bool {
	query, args, err := squirrel.Update("users").
		Set("codephone", code).
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		log.Println(messageError, err)
		return false
	}

	db := mysqlWorkerRepository.sessionFactory.GetDb()
	_, err = db.Exec(query, args...)
	if err != nil {
		log.Println(messageUpdate, err)
		return false
	}

	return true
}

func (mysqlWorkerRepository *MySQLWorkerRepository) UpdateUserPhoneVerification(code string) bool {
	query, args, err := squirrel.Select("verifiedphone", "email").
		From("users").
		Where(squirrel.Eq{"codephone": code}).
		ToSql()
	var verified bool
	var email string
	if err != nil {
		log.Println(messageError, err)
		return false
	}

	db := mysqlWorkerRepository.sessionFactory.GetDb()
	row := db.QueryRow(query, args...)
	err = row.Scan(&verified, &email)
	if err != nil {
		log.Println(messageBindData, err)
		return false
	}

	if verified {
		log.Println("User verified with phone, ", email)
		return false
	}

	query, args, err = squirrel.Update("users").
		Set("codephone", "").
		Set("verifiedphone", true).
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		log.Println(messageError, err)
		return false
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		log.Println(messageUpdate, err)
		return false
	}

	return true
}
