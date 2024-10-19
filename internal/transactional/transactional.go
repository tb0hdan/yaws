package transactional

import (
	"yaws/internal/transactional/sendgrid"
	"yaws/pkg/types"
)

type Transactional interface {
	Send(from, to types.Contact, subject string, message interface{}) error
}

func New() Transactional {
	return &sendgrid.SendGrid{}
}
