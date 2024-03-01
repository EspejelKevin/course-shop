package infrastructure

import (
	"os"
	"sync"
)

var settings *Settings
var lock = &sync.Mutex{}

type Settings struct {
	DriverName     string
	DataSourceName string
	Port           string
	Namespace      string
	APIVersion     string
	SecretKey      string
	TimeExpiration string
}

func NewSettings() *Settings {
	if settings == nil {
		lock.Lock()
		defer lock.Unlock()
		if settings == nil {
			settings = &Settings{
				DriverName:     os.Getenv("DRIVER_NAME"),
				DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
				Port:           os.Getenv("PORT"),
				Namespace:      os.Getenv("NAMESPACE"),
				APIVersion:     os.Getenv("VERSION_API"),
				SecretKey:      os.Getenv("SECRET_KEY"),
				TimeExpiration: os.Getenv("TIME_EXPIRATION"),
			}
		}
	}

	return settings
}
