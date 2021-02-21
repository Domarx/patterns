package main

import (
	"fmt"

	"github.com/domarx/patterns/behavioral/cor/handler"
)

func main() {
	Process(1, BuildChain(
		handler.NewAuthenticationHandler(),
		handler.NewAuthorizationHandler(),
		handler.NewRequestHandler(),
	))

	Process(2, BuildChain(
		handler.NewAuthenticationHandler(),
		handler.NewAuthenticationHandler(),
		handler.NewRequestHandler(),
	))

	Process(3, BuildChain(
		handler.NewAuthenticationHandler(),
		handler.NewRequestHandler(),
	))
	Process(4, BuildChain(
		handler.NewAuthorizationHandler(),
		handler.NewRequestHandler(),
	))
}

func BuildChain(handlers ...handler.Handler) handler.Handler {
	for i := 1; i < len(handlers); i++ {
		handlers[i-1].SetNext(handlers[i])
	}
	return handlers[0]
}

func Process(id int, chain handler.Handler) {
	fmt.Println(">>> processing request ", id, " <<< ")
	fmt.Println()
	fmt.Println(chain.Handle(handler.NewRequest(id)))
	fmt.Println()
}
