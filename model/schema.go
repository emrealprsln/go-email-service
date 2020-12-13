package model

import (
	"os"
)

const (
	RegistrationType  = 1
	ResetPasswordType = 2
)

type Mail interface {
	Send() error
}

type Schema interface {
	GetFrom() string
	GetTo() string
	GetReplyTo() string
	GetBody() Body
}

type schema struct {
	To   string
	Name string
	Type int
}

func NewSchema(to string, name string, emailType int) Schema {
	return &schema{
		To:   to,
		Name: name,
		Type: emailType,
	}
}

func (s schema) GetFrom() string {
	return os.Getenv("MAIL_FROM_ADDRESS")
}

func (s schema) GetTo() string {
	return s.To
}

func (s schema) GetReplyTo() string {
	return os.Getenv("MAIL_FROM_ADDRESS")
}

func (s schema) GetBody() Body {
	return s.getBodyInstance()
}

func (s schema) getBodyInstance() Body {
	if s.Type == RegistrationType {
		return NewRegistration(s.Name)
	}
	return NewResetPassword(s.Name)
}
