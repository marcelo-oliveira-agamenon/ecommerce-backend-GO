package utility

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

//SendEmailUtility function
func SendEmailUtility(toUser string, body string) bool {
	fromUser := GetDotEnv("EMAIL_PRO")
	userPassword := GetDotEnv("EMAIL_PRO_PASSWORD")

	mail := gomail.NewMessage()
	mail.SetHeader("From", fromUser)
	mail.SetHeader("To", toUser)
	mail.SetBody("text/html", body)

	send := gomail.NewDialer("smtp.gmail.com", 587, fromUser, userPassword)
	if err := send.DialAndSend(mail); err != nil {
		fmt.Print(err)
		return false
	  }

	return true
}