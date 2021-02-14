package company

type Company interface {
	Employee
	Add(component Employee)
	Remove(component Employee)
}

func NewCompany() Company {
	return &company{}
}

type company struct {
	employees []Employee
}

func (c *company) Execute() {
	for _, employee := range c.employees {
		employee.Execute()
	}
}

func (c *company) Add(component Employee) {
	c.employees = append(c.employees, component)
}

func (c *company) Remove(component Employee) {
	for i := 0; i < len(c.employees); i++ {
		if c.employees[i] == component {
			c.employees[i] = c.employees[len(c.employees)-1]
			c.employees = c.employees[:len(c.employees)-1]
		}
	}
}
