package handler

import (
	"encoding/json"
	"fmt"
)

func NewRequestHandler() Handler {
	return &requestHandler{}
}

type requestHandler struct{}

func (r requestHandler) SetNext(handler Handler) {}

func (r requestHandler) Handle(request Request) Response {
	fmt.Println("processing request No.: ", request.id)
	var errs = make([]error, 0)
	if !request.authenticated {
		errs = append(errs, fmt.Errorf("request %d not authenticated", request.id))
	}
	if !request.authorized {
		errs = append(errs, fmt.Errorf("request %d not authorized", request.id))
	}
	return Response{
		data: func() string {
			raw, _ := json.Marshal(map[string]interface {
			}{
				"id":            request.id,
				"authenticated": request.authenticated,
				"authorized":    request.authorized,
			})
			return string(raw)
		}(),
		err: errs,
	}
}
