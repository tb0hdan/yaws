package transactional

import (
	"yaws/internal/store/postgresql/models"
	"yaws/internal/transactional/sendgrid"
	"yaws/pkg/types"
)

const (
	SendGrid = "sendgrid"
)

type Transactional interface {
	Send(from, to types.Contact, subject string, message interface{}) error
}

type Sender struct {
	Transactional
	fromContact types.Contact
}

func (s *Sender) SendSimple(customer models.Customer, subject, message string) error {
	return s.Send(
		s.fromContact,
		types.Contact{Name: customer.Name, Email: customer.Email},
		subject,
		message,
	)
}

func New(senderType, apiKey string, fromContact types.Contact) *Sender {
	switch senderType {
	case SendGrid:
		return &Sender{
			Transactional: sendgrid.New(apiKey),
			fromContact:   fromContact,
		}
	default:
		return nil
	}
}
