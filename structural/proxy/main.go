package main

import (
	"fmt"

	"github.com/domarx/patterns/structural/proxy/structure"
)

func main() {
	service := structure.NewService()
	proxy := structure.NewProxyService(service)

	fmt.Println(proxy.Respond(structure.Request{
		ID:     1,
		Route:  "first-route",
		Params: "first-params",
	}))
	fmt.Println(proxy.Respond(structure.Request{
		ID:     1,
		Route:  "first-route",
		Params: "first-params",
	}))
	fmt.Println(proxy.Respond(structure.Request{
		ID:     2,
		Route:  "second-route",
		Params: "first-params",
	}))
}
