package gomail

import (
	"bytes"
	"errors"
	"html/template"
	"os"

	"github.com/ecommerce/ports"
	"gopkg.in/gomail.v2"
)

var (
	ErrorFileName   = errors.New("parsing file name")
	ErrorBufferExec = errors.New("buffer execution")
	ErrorSendEmail  = errors.New("email send failed")
)

type EmailDetails struct {
	fromUser     string
	userPassword string
	smtpConfig   string
}

func NewEmailService() ports.EmailService {
	return &EmailDetails{
		fromUser:     os.Getenv("EMAIL_PRO"),
		userPassword: os.Getenv("EMAIL_PRO_PASSWORD"),
		smtpConfig:   os.Getenv("EMAIL_PRO_SMTP"),
	}
}

func (em *EmailDetails) SendEmail(toUser string, fileName string, data interface{}, subject string) (bool, error) {
	fileEmail, err := template.ParseFiles(fileName)
	if err != nil {
		return false, ErrorFileName
	}
	buffer := new(bytes.Buffer)
	if err = fileEmail.Execute(buffer, data); err != nil {
		return false, ErrorBufferExec
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", em.fromUser)
	mail.SetHeader("To", toUser)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", buffer.String())

	send := gomail.NewDialer(em.smtpConfig, 587, em.fromUser, em.userPassword)
	if err := send.DialAndSend(mail); err != nil {
		return false, ErrorSendEmail
	}

	return true, nil
}
