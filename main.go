package main

import (
	"fmt"
	"os"

	"github.com/emrealprsln/go-email-service/driver"
	"github.com/emrealprsln/go-email-service/model"
	"github.com/joho/godotenv"
)

const (
	smtpDriver = "smtp"
	sesDriver  = "ses"
)

func main() {
	godotenv.Load()

	to := "to@domain"                    // Input - Http, queue, ...
	name := "Name"                       // Input - Http, queue, ...
	emailType := model.ResetPasswordType // Input - Http, queue, ...

	schema := model.NewSchema(to, name, emailType)
	mail := getMailInstance(schema)

	if err := mail.Send(); err != nil {
		fmt.Println(err)
	}
}

func getMailInstance(s model.Schema) model.Mail {
	d := os.Getenv("MAIL_DRIVER")

	if d == smtpDriver {
		return driver.NewSmtp(s)
	}
	if d == sesDriver {
		return driver.NewSes(s)
	}
	panic("unexpected driver")
}
