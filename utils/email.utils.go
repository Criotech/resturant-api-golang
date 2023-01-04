package utils

import (
	"net/smtp"
	"os"
)

func SendEmail(Subject string, Body string, To []string) error {
	from := os.Getenv("FROM_EMAIL_ADDR")
	password := os.Getenv("SMTP_PWD")

	to := To

	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	address := host + ":" + port

	subject := Subject
	body := Body
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	return err
}
