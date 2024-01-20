package ports

type EmailService interface {
	SendEmail(toUser string, fileName string, data interface{}, subject string) (bool, error)
}
