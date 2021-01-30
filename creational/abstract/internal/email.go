package internal

import "fmt"

type EmailFactory struct{}

func (e *EmailFactory) CreateMessenger() Messenger {
	return &email{}
}

type email struct {
	baseSender
}

func (e *email) SetRecipient(recipient string) Sender { e.baseSender.recipient = recipient; return e }

func (e *email) SetMessage(message []byte) Sender { e.baseSender.message = message; return e }

func (e *email) Send() error {
	fmt.Printf("Sending email to: %v, message = %v\n", e.recipient, string(e.message))
	return nil
}
