package remote

import "fmt"

func NewVolumeUpCommand(controller *Controller) Command {
	return &volumeUpCommand{controller}
}

type volumeUpCommand struct {
	*Controller
}

func (v *volumeUpCommand) Execute() error {
	fmt.Println("turning volume up")
	if v.volume+5 == 100 {
		return fmt.Errorf("volume at its max")
	}
	v.volume += 5
	return nil
}
