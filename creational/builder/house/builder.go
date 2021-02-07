package house

type builder interface {
	setFloor(floor string) builder
	setWalls(walls string) builder
	setRoof(roof string) builder
	setWindows(windows string) builder
	setPool(pool bool) builder
	setBBQGrill(grill bool) builder
	house() (House, error)
}

type baseBuilder struct {
	House
}

func (b *baseBuilder) setFloor(floor string) builder { b.Floor = floor; return b }

func (b *baseBuilder) setWalls(walls string) builder { b.Walls = walls; return b }

func (b *baseBuilder) setRoof(roof string) builder { b.Roof = roof; return b }

func (b *baseBuilder) setWindows(windows string) builder { b.Windows = windows; return b }

func (b *baseBuilder) setPool(pool bool) builder { b.Pool = pool; return b }

func (b *baseBuilder) setBBQGrill(grill bool) builder { b.BBQGrill = grill; return b }

func (b *baseBuilder) house() (House, error) { return b.House, nil }

type simpleHouseBuilder struct {
	builder builder
}

func (s simpleHouseBuilder) buildHouse() (House, error) {
	s.builder.
		setFloor("wood").
		setWalls("wood").
		setRoof("tiles").
		setWindows("plastic")
	return s.builder.house()
}

type richHouseBuilder struct {
	builder builder
}

func (r richHouseBuilder) buildHouse() (House, error) {
	r.builder.
		setFloor("marble").
		setWalls("brick").
		setWindows("wooden").
		setRoof("tiles").
		setPool(true).
		setBBQGrill(true)
	return r.builder.house()
}
