package utility

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

//SendEmailUtility function
func SendEmailUtility(toUser string, fileName string, data interface{}, subject string) bool {
	fromUser := GetDotEnv("EMAIL_PRO")
	userPassword := GetDotEnv("EMAIL_PRO_PASSWORD")
	smtpConfig := GetDotEnv("EMAIL_PRO_SMTP")

	fileEmail, err := template.ParseFiles(fileName)
	if err != nil {
		return false
	}
	buffer := new(bytes.Buffer)
	if err = fileEmail.Execute(buffer, data); err != nil {
		return false
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", fromUser)
	mail.SetHeader("To", toUser)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", buffer.String())

	send := gomail.NewDialer(smtpConfig, 587, fromUser, userPassword)
	if err := send.DialAndSend(mail); err != nil {
		return false
	}

	return true
}