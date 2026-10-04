package mail

import (
	"context"
	"fmt"
)

type Mailer interface {
	SendTicketEmail(
		ctx context.Context,
		email string,
		ticketID string,
	) error
}

// Mock
type SMTPMailer struct{}

func NewSMTPMailer() *SMTPMailer {

	return &SMTPMailer{}

}

func (m *SMTPMailer) SendTicketEmail(
	ctx context.Context,
	email string,
	ticketID string,
) error {

	fmt.Println(
		"Sending email to:",
		email,
		"ticket:",
		ticketID,
	)

	return nil

}
