package main

import (
	"context"
	"log"
	"net/smtp"

	"github.com/aws/aws-lambda-go/lambda"
)

type Employee struct {
	Name    string
	Email   string
	Year    string
	Salary  string
	IsLeave string
}

func handler(ctx context.Context) (string, error) {
	auth := smtp.PlainAuth("", "youremail@gmail.com", "youremail_password", "smtp.gmail.com")
	to := []string{"someone_email@gmail.com"}
	msg := []byte("To: someone_email@gmail.com\r\n" +
		"Subject: Test lambda\r\n" +
		"\r\n" +
		"this is test message from Lambda")
	err := smtp.SendMail("smtp.gmail.com:25", auth, "ledinhchung.it@gmail.com", to, msg)

	if err != nil {
		log.Fatal(err)
	}

	return "test", nil
}

func main() {
	lambda.Start(handler)
}
