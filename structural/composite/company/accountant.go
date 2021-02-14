package company

import (
	"fmt"
)

func NewAccountant(name string) Employee { return accountant(name) }

type accountant string

func (a accountant) Execute() { fmt.Printf("Accountant <%s> says hello\n", a) }
