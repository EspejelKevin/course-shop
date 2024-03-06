package infrastructure

import (
	"os"
	"sync"
)

var settings *Settings
var lockSettings = &sync.Mutex{}

type Settings struct {
	DriverName     string
	DataSourceName string
	Port           string
	Namespace      string
	APIVersion     string
	SecretKey      string
	TimeExpiration string
	SmtpHost       string
	SmtpUser       string
	SmtpPass       string
	SmtpPort       string
	EmailFrom      string
	PhoneFrom      string
}

func NewSettings() *Settings {
	if settings == nil {
		lockSettings.Lock()
		defer lockSettings.Unlock()
		if settings == nil {
			settings = &Settings{
				DriverName:     os.Getenv("DRIVER_NAME"),
				DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
				Port:           os.Getenv("PORT"),
				Namespace:      os.Getenv("NAMESPACE"),
				APIVersion:     os.Getenv("VERSION_API"),
				SecretKey:      os.Getenv("SECRET_KEY"),
				TimeExpiration: os.Getenv("TIME_EXPIRATION"),
				SmtpHost:       os.Getenv("SMTP_HOST"),
				SmtpUser:       os.Getenv("SMTP_USER"),
				SmtpPass:       os.Getenv("SMTP_PASS"),
				SmtpPort:       os.Getenv("SMTP_PORT"),
				EmailFrom:      os.Getenv("EMAIL_FROM"),
				PhoneFrom:      os.Getenv("PHONE_FROM"),
			}
		}
	}

	return settings
}
