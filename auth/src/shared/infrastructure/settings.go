package infrastructure

import "os"

type Settings struct {
	DriverName     string
	DataSourceName string
	Port           string `default:":3000"`
	Namespace      string
	APIVersion     string
}

func NewSettings() *Settings {
	return &Settings{
		DriverName:     os.Getenv("DRIVER_NAME"),
		DataSourceName: os.Getenv("DATA_SOURCE_NAME"),
		Port:           os.Getenv("PORT"),
		Namespace:      os.Getenv("NAMESPACE"),
		APIVersion:     os.Getenv("VERSION_API"),
	}
}
