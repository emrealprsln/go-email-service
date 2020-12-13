package model

import (
	"github.com/emrealprsln/go-email-service/util"
)

const (
	registrationSubject  = "Registration"
	registrationTemplate = "registration.html"
)

type registration struct {
	Name string
}

func NewRegistration(name string) Body {
	return &registration{
		Name: name,
	}
}

func (r registration) GetContent() string {
	return util.ParseTemplate(registrationTemplate, map[string]string{"name": r.Name})
}

func (r registration) GetSubject() string {
	return registrationSubject
}
