package main

import (
	"fmt"

	"github.com/domarx/patterns/behavioral/command/remote"
)

func main() {
	c := remote.NewController()
	printError(c.TurnUpVolume())
	printError(c.TurnDownVolume())
	printError(c.TurnDownVolume())
}

func printError(err error) {
	if err == nil {
		return
	}
	fmt.Println("ERROR > ", err)
}
