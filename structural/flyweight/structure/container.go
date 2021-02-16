package structure

/*
	Extrinsic state should be aggregated to the container object, which
	needs to contain a list for each state plus a list of references to the
	flyweight (intrinsic state) objects.
*/

// Container is an abstraction that the client can use to create context
// objects and use them without exposing implementation.
type Container interface {
	AddContext(x, y int, name, color, texture string)
	Draw()
}

func NewContainer() Container {
	return &container{}
}

type container struct {
	contexts []*context
}

// AddContext encapsulates the creation of intrinsic and extrinsic state objects.
// Intrinsic state object is returned from cache, whereas extrinsic state object is
// created on the spot.
func (c *container) AddContext(x, y int, name, color, texture string) {
	ctx := newContext(x, y, getFlyweight(name, color, texture))
	c.contexts = append(c.contexts, ctx)
}

func (c *container) Draw() {
	for _, ctx := range c.contexts {
		ctx.Draw()
	}
}
