package smtp

import (
	"github.com/MiteshSharma/project/model"
	gomail "gopkg.in/gomail.v2"
)

type Smtp struct {
}

func New() *Smtp {
	smtp := &Smtp{}
	return smtp
}

func (s Smtp) Send(to string, message model.NotificationMessage) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "email@gmail.com")
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", message.Title)
	if message.Type == "html" {
		mail.SetBody("text/html", message.Message)
	} else {
		mail.SetBody("text/plain", message.Message)
	}

	var dialer *gomail.Dialer
	dialer = gomail.NewDialer("smtp.gmail.com", 465, "username", "password")
	return dialer.DialAndSend(mail)
}
