package utility

import (
	"gopkg.in/gomail.v2"
)

//SendEmailUtility function
func SendEmailUtility(toUser string, body string, subject string) bool {
	fromUser := GetDotEnv("EMAIL_PRO")
	userPassword := GetDotEnv("EMAIL_PRO_PASSWORD")
	smtpConfig := GetDotEnv("EMAIL_PRO_SMTP")

	mail := gomail.NewMessage()
	mail.SetHeader("From", fromUser)
	mail.SetHeader("To", toUser)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	send := gomail.NewDialer(smtpConfig, 587, fromUser, userPassword)
	if err := send.DialAndSend(mail); err != nil {
		return false
	  }

	return true
}