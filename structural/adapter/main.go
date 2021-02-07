package main

import "github.com/domarx/patterns/structural/adapter/computer"

type client struct{}

func (c client) InsertRoundUSB(service computer.Service) {
	service.InsertRoundUSB()
}

func main() {
	c := client{}
	c.InsertRoundUSB(computer.RoundUSBMachine{})
	c.InsertRoundUSB(computer.SquareUSBMachineAdapter{Machine: computer.SquareUSBMachine{}})
}
