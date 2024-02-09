package infrastructure

import (
	"database/sql"
	"log"

	"sync"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once
var mysqlDatabase *MySQLDatabase

type MySQLDatabase struct {
	DB *sql.DB
}

func NewMySQLDatabase(driverName, dataSourceName string) *MySQLDatabase {
	once.Do(func() {
		DB, err := sql.Open(driverName, dataSourceName)

		if err != nil {
			log.Println("Failed to open database:", err)
		}

		mysqlDatabase = &MySQLDatabase{
			DB,
		}
	})

	return mysqlDatabase
}

func (mysqlDatabase *MySQLDatabase) IsUp() map[string]interface{} {
	data := map[string]interface{}{
		"status":  true,
		"message": "success",
	}
	query, args, err := squirrel.Select("1").ToSql()

	if err != nil {
		data["status"] = false
		data["message"] = "failed to create sql query"
		log.Println("Failed to create sql query:", err)
		return data
	}

	_, err = mysqlDatabase.DB.Exec(query, args...)

	if err != nil {
		data["status"] = false
		data["message"] = "failed to execute sql query"
		log.Println("Failed to execute sql query:", err)
		return data
	}

	return data
}
