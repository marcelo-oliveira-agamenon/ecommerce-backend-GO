package utility

import (
	"fmt"
	"net/smtp"
)

//SendEmailUtility function
func SendEmailUtility(toUser string) string {
	from := "marblack16@gmail.com"
	
	err := smtp.SendMail("smtp.gmail.com:587" ,smtp.PlainAuth("", from, "minhasenha1", "smtp.gmail.com"), from, []string{toUser}, []byte("dasdasd"))
	if err != nil {
		fmt.Print(err)
		return "erro"
	}

	return "sucesso"
}