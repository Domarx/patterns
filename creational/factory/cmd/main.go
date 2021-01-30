package main

import (
	"fmt"

	"github.com/domarx/patterns/creational/factory"
)

func main() {
	car, _ := factory.CreateVehicle(factory.Car)
	if err := car.Stop(); err != nil {
		fmt.Println(err)
	}
	if err := car.Drive(); err != nil {
		fmt.Println(err)
	}
	if err := car.Drive(); err != nil {
		fmt.Println(err)
	}
	if err := car.Stop(); err != nil {
		fmt.Println(err)
	}

	boat, _ := factory.CreateVehicle(factory.Boat)
	if err := boat.Stop(); err != nil {
		fmt.Println(err)
	}
	if err := boat.Drive(); err != nil {
		fmt.Println(err)
	}
	if err := boat.Drive(); err != nil {
		fmt.Println(err)
	}
	if err := boat.Stop(); err != nil {
		fmt.Println(err)
	}
}
