package company

import (
	"fmt"
)

func NewDepartment(name string) Company { return &department{name: name} }

type department struct {
	name      string
	employees []Employee
}

func (d department) Execute() {
	fmt.Printf("====== Department <%s> =======\n", d.name)
	for _, e := range d.employees {
		e.Execute()
	}
}

func (d *department) Add(component Employee) {
	d.employees = append(d.employees, component)
}

func (d *department) Remove(component Employee) {
	for i, e := range d.employees {
		if e == component {
			d.employees[i] = d.employees[len(d.employees)-1]
			d.employees = d.employees[:len(d.employees)-1]
		}
	}
}
