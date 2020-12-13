package driver

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/emrealprsln/go-email-service/model"
)

type smtpConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func getConfig() smtpConfig {
	return smtpConfig{
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	}
}

func (s smtpConfig) getAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type smtpServer struct {
	Schema model.Schema
	Config smtpConfig
}

func NewSmtp(s model.Schema) model.Mail {
	return &smtpServer{
		Schema: s,
		Config: getConfig(),
	}
}

func (s smtpServer) Send() error {
	auth := smtp.PlainAuth("", s.Config.Username, s.Config.Password, s.Config.Host)
	to := s.Schema.GetTo()

	header := make(map[string]string)
	header["From"] = s.Schema.GetFrom()
	header["To"] = to
	header["Subject"] = s.Schema.GetBody().GetSubject()
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""

	body := ""
	content := s.Schema.GetBody().GetContent()
	for k, v := range header {
		body += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	body += fmt.Sprintf("\r\n%s", content)

	err := smtp.SendMail(s.Config.getAddress(), auth, s.Schema.GetFrom(), []string{to}, []byte(body))
	if err != nil {
		return err
	}
	return nil
}
