package gomail

import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"text/template"

	"github.com/ecommerce/ports"
	"gopkg.in/gomail.v2"
)

var (
	ErrorSendEmail = errors.New("email send failed")
)

type EmailDetails struct {
	fromUser     string
	userPassword string
	smtpConfig   string
	smtpPort     string
}

func NewEmailService() ports.EmailService {
	return &EmailDetails{
		fromUser:     os.Getenv("EMAIL_PRO"),
		userPassword: os.Getenv("EMAIL_PRO_PASSWORD"),
		smtpConfig:   os.Getenv("EMAIL_PRO_SMTP"),
		smtpPort:     os.Getenv("EMAIL_PRO_PORT"),
	}
}

func (em *EmailDetails) SendEmail(toUser string, fileName string, data interface{}, subject string) (bool, error) {
	fileEmail, err := template.ParseFiles(fileName)
	if err != nil {
		return false, ErrorSendEmail
	}
	buffer := new(bytes.Buffer)
	if err = fileEmail.Execute(buffer, data); err != nil {
		return false, ErrorSendEmail
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", em.fromUser)
	mail.SetHeader("To", toUser)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", buffer.String())

	port, errP := strconv.Atoi(em.smtpPort)
	if errP != nil {
		return false, ErrorSendEmail
	}
	send := gomail.NewDialer(em.smtpConfig, port, em.fromUser, em.userPassword)
	if err := send.DialAndSend(mail); err != nil {
		return false, ErrorSendEmail
	}

	return true, nil
}
