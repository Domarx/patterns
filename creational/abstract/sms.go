package abstract

import "fmt"

type smsFactory struct{}

func (s *smsFactory) CreateMessenger() Messenger {
	return &sms{}
}

type sms struct {
	baseSender
}

func (s *sms) SetRecipient(recipient string) Sender { s.baseSender.recipient = recipient; return s }

func (s *sms) SetMessage(message []byte) Sender { s.baseSender.message = message; return s }

func (s *sms) Send() error {
	fmt.Printf("Sending sms to: %v, message = %v\n", s.recipient, string(s.message))
	return nil
}
