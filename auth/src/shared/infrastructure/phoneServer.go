package infrastructure

import (
	"log"
	"sync"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var phoneServer *PhoneServer
var lockPhoneServer = &sync.Mutex{}

type PhoneServer struct {
	client *twilio.RestClient
}

func NewPhoneServer() *PhoneServer {
	if phoneServer == nil {
		lockPhoneServer.Lock()
		defer lockPhoneServer.Unlock()
		if phoneServer == nil {
			client := twilio.NewRestClient()
			phoneServer = &PhoneServer{client}
		}
	}
	return phoneServer
}

func (phoneServer *PhoneServer) IsUp() map[string]interface{} {
	data := map[string]interface{}{
		"status":  true,
		"message": "success",
	}

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+5217731088261")
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("**Success**")

	_, err := phoneServer.client.Api.CreateMessage(params)

	if err != nil {
		data["status"] = false
		data["message"] = "failed to send test message"
		log.Println("Failed to send message:", err)
		return data
	}

	return data
}

func (phoneServer *PhoneServer) GetClient() *twilio.RestClient {
	return phoneServer.client
}
