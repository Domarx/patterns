package remote

import "fmt"

func NewVolumeDownCommand(controller *Controller) Command {
	return &volumeDownCommand{controller}
}

type volumeDownCommand struct {
	*Controller
}

func (v volumeDownCommand) Execute() error {
	fmt.Println("turning volume down")
	if v.volume-5 < 0 {
		return fmt.Errorf("volume at it minimum")
	}
	v.volume -= 5
	return nil
}
