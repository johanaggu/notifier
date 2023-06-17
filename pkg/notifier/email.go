package notifier

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/appengine/mail"
)

type email struct{}

func NewEmail() *email {
	return &email{}
}

func (e *email) Send(from, to, message string) error {
	ctx := context.Background()
	addr := to

	msg := &mail.Message{
		Sender:  "Example.com Support <support@example.com>",
		To:      []string{addr},
		Subject: "Confirm your registration",
		Body:    fmt.Sprintf(confirmMessage, message),
	}
	if err := mail.Send(ctx, msg); err != nil {
		log.Printf("error sending message - %s:%s", ErrSendMessage, err)
		return fmt.Errorf("error sending message - %s", ErrSendMessage)
	}

	return nil
}

const confirmMessage = `
Thank you for creating an account!
Please confirm your email address by clicking on the link below:

%s
`
