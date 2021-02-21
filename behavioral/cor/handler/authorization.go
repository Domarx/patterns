package handler

import "fmt"

func NewAuthorizationHandler() Handler {
	return &authorizationHandler{}
}

type authorizationHandler struct {
	next Handler
}

func (a *authorizationHandler) SetNext(handler Handler) {
	a.next = handler
}

func (a *authorizationHandler) Handle(request Request) Response {
	fmt.Println("authorizing request No.: ", request.id)
	if request.authorized {
		fmt.Println("request authorized already")
		return a.next.Handle(request)
	}
	request.authorized = true
	return a.next.Handle(request)
}
