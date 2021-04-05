package main

import "fmt"

type Sender interface {
	Send(message string)
}

type smsSender struct{}

func (sms *smsSender) Send(message string) { fmt.Println("sending SMS: ", message) }

type emailSender struct{}

func (email *emailSender) Send(message string) { fmt.Println("sending EMAIL: ", message) }

type Context struct {
	sender Sender
}

func (c *Context) SetSender(sender Sender) { c.sender = sender }

func (c *Context) Do(message string) {
	c.sender.Send(message)
}

func Send(context *Context, message string, senderType int) {
	if senderType&1 == 0 {
		context.SetSender(&smsSender{})
		context.Do(message)
	} else {
		context.SetSender(&emailSender{})
		context.Do(message)
	}
}

func main() {
	c := Context{}
	for i := 0; i < 4; i++ {
		Send(&c, "Hello world!", i)
	}
}
