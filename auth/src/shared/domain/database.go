package domain

import "database/sql"

type Database interface {
	IsUp() map[string]interface{}
	GetDb() *sql.DB
}
