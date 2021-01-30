package abstract

import "fmt"

type emailFactory struct {
	email email
}

func (e *emailFactory) CreateMessenger() Messenger {
	return &e.email
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
