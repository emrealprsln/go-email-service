package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/emrealprsln/go-email-service/util"
)

const (
	resetPasswordSubject  = "Reset Password"
	resetPasswordTemplate = "reset_password.html"
)

type Body interface {
	GetSubject() string
	GetContent() string
}

type resetPassword struct {
	Name string
}

func NewResetPassword(name string) Body {
	return &resetPassword{
		Name: name,
	}
}

func (r resetPassword) GetContent() string {
	return util.ParseTemplate(resetPasswordTemplate, map[string]string{"name": r.Name, "link": r.generateLink()})
}

func (r resetPassword) GetSubject() string {
	return resetPasswordSubject
}

func (r resetPassword) generateLink() string {
	hash := md5.Sum([]byte(r.Name))
	return fmt.Sprintf("%s/reset?token=%s", os.Getenv("BASE_URL"), hex.EncodeToString(hash[:]))
}
