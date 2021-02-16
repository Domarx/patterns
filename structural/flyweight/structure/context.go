package structure

// _ structure contains both mutable and immutable states, some of which
// are unique and others are shared amongst the same type of objects.
type _ struct {
	// Unique (mutable) values
	x, y int

	// Immutable values that are shared amongst many objects. Assuming these
	// values are costing lots of memory, such object should be split into two
	// objects where one contains mutable values, and other contains memory costly
	// constant and shared state amongst the same type of objects.
	name    string
	color   string
	texture string
}

// Context should contain extrinsic (mutable) values that are unique
// to each instance of the Context. Context state can be altered by
// other objects.
type context struct {
	x, y      int
	flyweight *flyweight
}

// Draw is simply an action that a context object can perform
func (c *context) Draw() {
	c.flyweight.Draw(c.x, c.y)
}

func newContext(x, y int, flyweight *flyweight) *context {
	return &context{
		x:         x,
		y:         y,
		flyweight: flyweight,
	}
}
