package main

import "github.com/domarx/patterns/creational/prototype/shape"

func main() {
	cache := shape.NewPrototypeCachce()
	circleClone := cache.MustGetPrototype("circle")
	squareClone := cache.MustGetPrototype("square")

	DrawShapes(circleClone, squareClone)
}

func DrawShapes(drawables ...shape.Shape) {
	for _, drawable := range drawables {
		drawable.Draw()
	}
}
