package structure

type Service interface {
	Respond(request Request) Response
}

type Request struct {
	ID     int
	Route  string
	Params string
}

type Response struct {
	Data string
}
