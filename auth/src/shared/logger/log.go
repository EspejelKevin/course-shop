package logger

import (
	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
)

type Log struct {
	logger    *logrus.Logger
	TracingId string
}

func NewLogger() *Log {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})
	return &Log{logger, ""}
}

func (log *Log) Info(origin, service, message string, measurement *Measurement) {
	var measurementToMap map[string]interface{} = map[string]interface{}{}

	if measurement != nil {
		measurementToMap = structs.Map(measurement)
	}

	log.logger.WithFields(logrus.Fields{
		"origin":      origin,
		"service":     service,
		"measurement": measurementToMap,
	}).Info(message)
}

func (log *Log) Error(origin, service, message, err string, measurement Measurement) {
	log.logger.SetReportCaller(true)
	log.logger.WithFields(logrus.Fields{
		"origin":      origin,
		"service":     service,
		"error":       err,
		"measurement": structs.Map(measurement),
	}).Error(message)
}
