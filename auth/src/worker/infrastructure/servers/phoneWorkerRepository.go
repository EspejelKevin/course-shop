package servers

import (
	"auth/src/shared/domain"
	"auth/src/worker/domain/entities"
	"log"
	"sync"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var phoneWorkerRepository *PhoneWorkerRepository
var lockPhoneServer = &sync.Mutex{}

type PhoneWorkerRepository struct {
	client domain.Phone
}

func NewPhoneWorkerRepository(client domain.Phone) *PhoneWorkerRepository {
	if phoneWorkerRepository == nil {
		lockPhoneServer.Lock()
		defer lockPhoneServer.Unlock()
		if phoneWorkerRepository == nil {
			phoneWorkerRepository = &PhoneWorkerRepository{
				client,
			}
		}
	}
	return phoneWorkerRepository
}

func (phoneWorkerRepository *PhoneWorkerRepository) IsUp() bool {
	data := phoneWorkerRepository.client.IsUp()
	status := data["status"].(bool)
	message := data["message"].(string)

	if status {
		log.Println("Twilio Server is up", message)
	} else {
		log.Println("Twilio Server is down", message)
	}

	return status
}

func (phoneWorkerRepository *PhoneWorkerRepository) SendMessage(message *entities.Message) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(message.To)
	params.SetFrom(message.From)
	params.SetBody(message.Body)

	client := phoneWorkerRepository.client.GetClient()

	_, err := client.Api.CreateMessage(params)

	if err != nil {
		log.Println("Failed to send message:", err)
	}

	return err
}
