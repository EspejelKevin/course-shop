package infrastructure

import (
	"log"
	"strconv"
	"sync"

	"gopkg.in/gomail.v2"
)

var mailServer *MailServer
var lockMailServer = &sync.Mutex{}

type MailServer struct {
	SMTPServer *gomail.Dialer
}

func NewMailServer(SmtpHost, SmtpUser, SmtpPass, SmtpPort string) *MailServer {
	port, _ := strconv.Atoi(SmtpPort)
	if mailServer == nil {
		lockMailServer.Lock()
		defer lockMailServer.Unlock()
		if mailServer == nil {
			SMTPServer := gomail.NewDialer(SmtpHost, port, SmtpUser, SmtpPass)
			mailServer = &MailServer{
				SMTPServer,
			}
		}
	}

	return mailServer
}

func (mailServer *MailServer) IsUp() map[string]interface{} {
	data := map[string]interface{}{
		"status":  true,
		"message": "success",
	}

	message := gomail.NewMessage()
	message.SetHeader("From", "courseshopa@gmail.com")
	message.SetHeader("To", "kevin.espejelmtz@gmail.com")
	message.SetHeader("Subject", "Test running SMTP server")
	message.SetBody("text/html", "<h1>Test running</h1>")

	if err := mailServer.SMTPServer.DialAndSend(message); err != nil {
		data["status"] = false
		data["message"] = "failed to create mail"
		log.Println("Failed to send mail:", err)
		return data
	}

	return data
}

func (mailServer *MailServer) GetSMTP() *gomail.Dialer {
	return mailServer.SMTPServer
}
