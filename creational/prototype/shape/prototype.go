package shape

import "fmt"

func NewPrototypeCachce() PrototypeCache {
	return map[string]Prototype{
		"circle": &circle{
			kind:   "circle",
			radius: 3.14,
		},
		"square": &square{
			x:    2.5,
			y:    2.5,
			kind: "square",
		},
	}

}

type PrototypeCache map[string]Prototype

func (p PrototypeCache) MustGetPrototype(_type string) Prototype {
	prototype, found := p[_type]
	if !found {
		panic("prototype not found")
	}
	return prototype.Clone()
}

type Shape interface {
	Draw()
}

type Prototype interface {
	Shape
	Clone() Prototype
}

type square struct {
	x, y float64
	kind string
}

func (s square) Draw() {
	fmt.Printf("Drawing %s: x = %.2f, y = %.2f\n", s.kind, s.x, s.y)
}

func (s square) Clone() Prototype {
	return &square{
		x:    s.x,
		y:    s.y,
		kind: s.kind + "_clone",
	}
}

type circle struct {
	kind   string
	radius float64
}

func (c circle) Draw() {
	fmt.Printf("Drawing %s: radius = %.2f\n", c.kind, c.radius)
}

func (c circle) Clone() Prototype {
	return &circle{
		kind:   c.kind + "_clone",
		radius: c.radius,
	}
}
