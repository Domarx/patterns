package factory

import "fmt"

type car struct {
	moving bool
}

func (c *car) Drive() error {
	if c.moving {
		return fmt.Errorf("car is already moving")
	}
	c.moving = true
	fmt.Println("car is moving")
	return nil
}

func (c *car) Stop() error {
	if !c.moving {
		return fmt.Errorf("car cannot stop because it is not moving")
	}
	c.moving = false
	fmt.Println("car stopped")
	return nil
}
