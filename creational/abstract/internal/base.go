package internal

type Sender interface {
	SetRecipient(recipient string) Sender
	SetMessage(message []byte) Sender
}

type Messenger interface {
	Sender
	Send() error
}

type baseSender struct {
	recipient string
	message   []byte
}
