package servers

import (
	"auth/src/shared/domain"
	"auth/src/shared/logger"
	"auth/src/worker/domain/entities"
	"fmt"
	"log"
	"sync"
	"time"

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

func (phoneWorkerRepository *PhoneWorkerRepository) IsUp(log *logger.Log) bool {
	start := time.Now()
	data := phoneWorkerRepository.client.IsUp()
	status := data["status"].(bool)
	message := data["message"].(string)
	timeElapsed := fmt.Sprint(time.Since(start).Milliseconds()) + "ms"
	measurement := logger.Measurement{TimeElapsed: timeElapsed, Object: map[string]interface{}{}}

	if status {
		log.Info("External", "Twilio", message, &measurement)
	} else {
		log.Error("External", "Twilio", "Twilio is down", message, &measurement)
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
