package main

import "github.com/domarx/patterns/structural/flyweight/structure"

func main() {
	container := structure.NewContainer()

	for i := 0; i < 3; i++ {
		container.AddContext(i*1, i*1, "bullet", "gold", "smooth")
		container.AddContext(i*2, i*2, "shrapnel", "gray", "sharp")
		container.AddContext(i*3, i*3, "metal pieces", "dark gray", "rough")
	}
	container.Draw()
}
