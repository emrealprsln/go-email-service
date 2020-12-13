package driver

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/emrealprsln/go-email-service/model"
)

type sesConfig struct {
	Region          string
	AccessKeyId     string
	SecretAccessKey string
}

type sesServer struct {
	Schema model.Schema
	Config sesConfig
}

func getSesConfig() sesConfig {
	return sesConfig{
		Region:          os.Getenv("AWS_DEFAULT_REGION"),
		AccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
}

func NewSes(s model.Schema) model.Mail {
	return &sesServer{
		Schema: s,
		Config: getSesConfig(),
	}
}

func (s sesServer) newSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s.Config.Region),
		Credentials: credentials.NewStaticCredentials(s.Config.AccessKeyId, s.Config.SecretAccessKey, ""),
	})
	if err != nil {
		fmt.Println("session creating error,", err)
	}
	return sess
}

func (s sesServer) Send() error {
	sesSession := ses.New(s.newSession())
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(s.Schema.GetTo()),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data:    aws.String(s.Schema.GetBody().GetContent()),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String(s.Schema.GetBody().GetSubject()),
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(s.Schema.GetFrom()),
		ReplyToAddresses: []*string{
			aws.String(s.Schema.GetReplyTo()),
		},
	}

	if _, err := sesSession.SendEmail(input); err != nil {
		return err
	}
	return nil
}
