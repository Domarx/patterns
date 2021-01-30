package singleton

import (
	"fmt"
	"sync"
)

// Greeter prints greeting message
type Greeter interface {
	Greet(name string)
}

var mux = new(sync.Mutex)

type greeter struct {
	id int
}

//Greet prints personal greet message
func (g greeter) Greet(name string) {
	fmt.Printf("Greeter %d says: Hello, %s", g.id, name)
}

var instance *greeter

//Create creates a singleton of Greeter
func Create() Greeter {
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = &greeter{id: 0}
		}
	}
	return instance
}
