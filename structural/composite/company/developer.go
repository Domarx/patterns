package company

import (
	"fmt"
)

func NewDeveloper(name string) Employee { return developer(name) }

type developer string

func (d developer) Execute() { fmt.Printf("Developer <%s> says hello\n", d) }
