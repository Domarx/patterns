package remote

func NewController() *Controller {
	return &Controller{}
}

type Controller struct {
	volume int
}

func (c *Controller) TurnUpVolume() error {
	return NewVolumeUpCommand(c).Execute()
}

func (c *Controller) TurnDownVolume() error {
	return NewVolumeDownCommand(c).Execute()
}
