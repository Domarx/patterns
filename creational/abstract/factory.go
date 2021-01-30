package abstract

import (
	"fmt"

	"github.com/domarx/patterns/creational/abstract/internal"
)

type Messenger = internal.Messenger

type senderType int

const (
	Email senderType = iota
	SMS
)

func NewSenderFactory(sender senderType) (factory MessengerFactory, err error) {
	switch sender {
	case Email:
		factory = &internal.EmailFactory{}
	case SMS:
		factory = &internal.SMSFactory{}
	default:
		err = fmt.Errorf("no sender of such type %d", sender)
	}
	return factory, err
}

type MessengerFactory interface {
	CreateMessenger() Messenger
}
