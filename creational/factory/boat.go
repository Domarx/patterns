package factory

import "fmt"

type boat struct {
	moving bool
}

func (b *boat) Drive() error {
	if b.moving {
		return fmt.Errorf("boat is already moving")
	}
	b.moving = true
	fmt.Println("boat is moving")
	return nil
}

func (b *boat) Stop() error {
	if !b.moving {
		return fmt.Errorf("boat cannot stop because boat is not moving")
	}
	b.moving = false
	fmt.Println("boat stopped")
	return nil
}
