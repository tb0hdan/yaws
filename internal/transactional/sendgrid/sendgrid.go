package sendgrid

import (
	"fmt"
	"log"

	"yaws/pkg/types"
	"yaws/pkg/xerrors"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	MessageNotAString = xerrors.XError("message is not a string")
)

type SendGrid struct {
	APIKey string
}

func (s SendGrid) Send(from, to types.Contact, subject string, message interface{}) error {
	mailFrom := mail.NewEmail(from.Name, from.Email)
	mailTo := mail.NewEmail(to.Name, to.Email)
	htmlContent := fmt.Sprintf("<strong>%s</strong>", message)
	msg, ok := message.(string)
	if !ok {
		return MessageNotAString
	}
	mailMessage := mail.NewSingleEmail(mailFrom, subject, mailTo, msg, htmlContent)
	client := sendgrid.NewSendClient(s.APIKey)
	response, err := client.Send(mailMessage)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
	return err
}

func New(apiKey string) SendGrid {
	return SendGrid{APIKey: apiKey}
}
