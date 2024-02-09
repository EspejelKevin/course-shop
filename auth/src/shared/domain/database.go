package domain

type Database interface {
	IsUp() map[string]interface{}
}
