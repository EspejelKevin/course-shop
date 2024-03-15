package domain

import "gopkg.in/gomail.v2"

type Mail interface {
	IsUp() map[string]interface{}
	GetSMTP() *gomail.Dialer
}
