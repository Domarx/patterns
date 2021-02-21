package handler

import "fmt"

func NewAuthenticationHandler() Handler {
	return &authenticationHandler{}
}

type authenticationHandler struct {
	next Handler
}

func (a *authenticationHandler) SetNext(handler Handler) {
	a.next = handler
}

func (a *authenticationHandler) Handle(request Request) Response {
	fmt.Println("authenticating request No.: ", request.id)
	if request.authenticated {
		fmt.Println("request authenticated already")
		return a.next.Handle(request)
	}
	request.authenticated = true
	return a.next.Handle(request)
}
