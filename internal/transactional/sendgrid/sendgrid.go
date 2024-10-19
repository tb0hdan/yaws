package sendgrid

import (
	"fmt"
	"log"
	"os"

	"yaws/pkg/types"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
}

func (s SendGrid) Send(from, to types.Contact, subject string, message interface{}) error {
	mailFrom := mail.NewEmail(from.Name, from.Email)
	mailTo := mail.NewEmail(from.Name, from.Email)
	htmlContent := fmt.Sprintf("<strong>%s</strong>", message)
	mailMessage := mail.NewSingleEmail(mailFrom, subject, mailTo, message.(string), htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
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
