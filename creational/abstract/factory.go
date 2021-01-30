package abstract

import "fmt"

type senderType int

const (
	Email senderType = iota
	SMS
)

func NewSenderFactory(sender senderType) (factory MessengerFactory, err error) {
	switch sender {
	case Email:
		factory = &emailFactory{}
	case SMS:
		factory = &smsFactory{}
	default:
		err = fmt.Errorf("no sender of such type %d", sender)
	}
	return factory, err
}

type MessengerFactory interface {
	CreateMessenger() Messenger
}

type Sender interface {
	SetRecipient(recipient string) Sender
	SetMessage(message []byte) Sender
}

type Messenger interface {
	Sender
	Send() error
}
