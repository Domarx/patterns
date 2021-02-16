package structure

import "fmt"

func NewProxyService(service Service) Service {
	return &proxy{
		service: service,
		cache:   make(map[Request]Response),
	}
}

type proxy struct {
	service Service
	cache   map[Request]Response
}

func (p proxy) Respond(request Request) Response {
	if response, found := p.cache[request]; found {
		fmt.Printf("Returning from cache: %v -> %v\n", request, response)
		return response
	}
	fmt.Println("Calling original service")
	response := p.service.Respond(request)
	p.cache[request] = response
	return response
}
