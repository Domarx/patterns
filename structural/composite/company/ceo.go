package company

import (
	"fmt"
)

func NewCEO(name string) Employee { return ceo(name) }

type ceo string

func (c ceo) Execute() { fmt.Printf("CEO <%s> says hello\n", c) }
