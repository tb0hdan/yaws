package transactional

import (
	"yaws/internal/transactional/sendgrid"
	"yaws/pkg/types"
)

const (
	SendGrid = "sendgrid"
)

type Transactional interface {
	Send(from, to types.Contact, subject string, message interface{}) error
}

func New(senderType, apiKey string) Transactional {
	switch senderType {
	case SendGrid:
		return sendgrid.New(apiKey)
	default:
		return nil
	}
}
