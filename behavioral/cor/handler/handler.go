package handler

func NewRequest(id int) Request {
	return Request{id: id}
}

type Request struct {
	id            int
	authenticated bool
	authorized    bool
}

type Response struct {
	data string
	err  []error
}

func (r Response) Error() string {
	var out string
	for _, e := range r.err {
		out += e.Error()
	}
	return "ERROR >: " + out
}

type Handler interface {
	SetNext(handler Handler)
	Handle(request Request) Response
}
