package structure

import "fmt"

func NewService() Service {
	return &service{}
}

type service struct{}

func (s service) Respond(request Request) Response {
	response := fmt.Sprintf("RESPONSE TO: %d %s %s", request.ID, request.Params, request.Route)
	return Response{Data: response}
}
