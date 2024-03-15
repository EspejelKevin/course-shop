package domain

import "github.com/twilio/twilio-go"

type Phone interface {
	IsUp() map[string]interface{}
	GetClient() *twilio.RestClient
}
