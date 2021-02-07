package computer

import "fmt"

type Service interface {
	InsertRoundUSB()
}

type RoundUSBMachine struct{}

func (r RoundUSBMachine) InsertRoundUSB() {
	fmt.Println("round USB Machine inserts round usb")
}

type SquareUSBMachineAdapter struct {
	Machine SquareUSBMachine
}

func (s SquareUSBMachineAdapter) InsertRoundUSB() {
	s.Machine.InsertSquareUSB()
}

type SquareUSBMachine struct{}

func (s SquareUSBMachine) InsertSquareUSB() {
	fmt.Println("square USB Machine inserts square usb")
}
