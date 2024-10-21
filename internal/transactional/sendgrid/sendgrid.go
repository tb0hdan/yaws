package sendgrid

import (
	"fmt"
	"log"

	"yaws/pkg/types"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	APIKey string
}

func (s SendGrid) Send(from, to types.Contact, subject string, message interface{}) error {
	mailFrom := mail.NewEmail(from.Name, from.Email)
	mailTo := mail.NewEmail(to.Name, to.Email)
	htmlContent := fmt.Sprintf("<strong>%s</strong>", message)
	mailMessage := mail.NewSingleEmail(mailFrom, subject, mailTo, message.(string), htmlContent)
	client := sendgrid.NewSendClient(s.APIKey)
	response, err := client.Send(mailMessage)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return err
}

func New(apiKey string) SendGrid {
	return SendGrid{APIKey: apiKey}
}
