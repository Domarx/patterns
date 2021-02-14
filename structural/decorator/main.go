package main

import (
	"fmt"
)

type Service interface {
	DoStuff()
}

type service struct {
	name string
}

func (s service) DoStuff() { fmt.Printf("Service %s called\n", s.name) }

type decorator1 struct {
	name    string
	service Service
}

func (d decorator1) DoStuff() {
	fmt.Printf("Decorator %s: BEFORE\n", d.name)
	d.service.DoStuff()
	fmt.Printf("Decorator %s: AFTER\n", d.name)
}

type decorator2 struct {
	name    string
	service Service
}

func (d decorator2) DoStuff() {
	fmt.Printf("Decorator %s: BEFORE\n", d.name)
	d.service.DoStuff()
	fmt.Printf("Decorator %s: AFTER\n", d.name)
}

func main() {
	base := &service{name: "ORIGINAL"}
	d1 := &decorator1{name: "FIRST", service: base}
	d2 := &decorator2{name: "SECOND", service: d1}

	base.DoStuff()
	fmt.Println()

	d1.DoStuff()
	fmt.Println()

	d2.DoStuff()
}
