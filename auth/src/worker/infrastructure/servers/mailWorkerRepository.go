package servers

import (
	"auth/src/shared/domain"
	"auth/src/shared/utils"
	"auth/src/worker/domain/entities"
	"bytes"
	"log"
	"sync"

	"gopkg.in/gomail.v2"
)

var mailWorkerRepository *MailWorkerRepository
var lockMail = &sync.Mutex{}

type MailWorkerRepository struct {
	mail domain.Mail
}

func NewMailWorkerRepository(mail domain.Mail) *MailWorkerRepository {
	if mailWorkerRepository == nil {
		lockMail.Lock()
		defer lockMail.Unlock()
		if mailWorkerRepository == nil {
			mailWorkerRepository = &MailWorkerRepository{
				mail,
			}
		}
	}
	return mailWorkerRepository
}

func (mailWorkerRepository *MailWorkerRepository) IsUp() bool {
	data := mailWorkerRepository.mail.IsUp()
	status := data["status"].(bool)
	message := data["message"].(string)

	if status {
		log.Println("SMTP is up", message)
	} else {
		log.Println("SMTP is down", message)
	}

	return status
}

func (mailWorkerRepository *MailWorkerRepository) SendMail(mail *entities.Email) error {
	var body bytes.Buffer

	template, err := utils.ParseTemplateDir("src/worker/domain/templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verification-code.html", &mail)

	message := gomail.NewMessage()
	message.SetHeader("From", mail.From)
	message.SetHeader("To", mail.To)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody("text/html", body.String())

	smtpClient := mailWorkerRepository.mail.GetSMTP()

	if err := smtpClient.DialAndSend(message); err != nil {
		log.Println("Failed to send mail:", err)
	}

	return err
}
