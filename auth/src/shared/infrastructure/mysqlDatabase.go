package infrastructure

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatabase struct {
	DB *sql.DB
}

func NewMySQLDatabase(driverName, dataSourceName string) *MySQLDatabase {
	DB, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Println("Failed to open database:", err)
		return nil
	}

	return &MySQLDatabase{
		DB,
	}
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

func (mysqlDatabase *MySQLDatabase) Close() error {
	err := mysqlDatabase.DB.Close()

	if err != nil {
		log.Println("Error closing the connection:", err)
	}

	return err
}
