package main

import (
	"github.com/domarx/patterns/creational/abstract"
)

func main() {
	factory, _ := abstract.NewSenderFactory(abstract.SMS)
	messenger := factory.CreateMessenger()
	messenger.
		SetRecipient("John Doe").
		SetMessage([]byte(`"Hello, World!"`))
	_ = messenger.Send()

	factory, _ = abstract.NewSenderFactory(abstract.Email)
	messenger = factory.CreateMessenger()
	messenger.SetRecipient("John Doe").SetMessage([]byte(`"Good Bye, World!"`))
	_ = messenger.Send()

}
