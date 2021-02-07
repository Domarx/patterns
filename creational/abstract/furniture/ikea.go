package furniture

type ikeaFactory struct{}

func (i ikeaFactory) GetFactory(factory string) Factory {
	var f Factory
	switch factory {
	case "modern":
		f = &modernIkeaFactory{}
	case "classic":
		f = &classicIkeaFactory{}
	}
	return f
}

type classicIkeaFactory struct{}

func (c classicIkeaFactory) NewChair() Chair {
	return &classicIkeaChair{
		details: ChairDetails{
			Manufacturer: "ikea",
			Type:         "classic",
			LegCount:     4,
			Material:     "wood",
			Color:        "brown",
		},
		price: 99.99,
	}
}

func (c classicIkeaFactory) NewTable() Table {
	return &classicIkeaTable{
		details: TableDetails{
			Manufacturer: "ikea",
			Type:         "classic",
			LegCount:     4,
			Material:     "wood",
			Color:        "brown",
		},
		price: 199.99,
	}
}

type modernIkeaFactory struct{}

func (m modernIkeaFactory) NewChair() Chair {
	return &modernIkeaChair{
		details: ChairDetails{
			Manufacturer: "ikea",
			Type:         "modern",
			LegCount:     3,
			Material:     "leather",
			Color:        "brown",
		},
		price: 194.99,
	}
}

func (m modernIkeaFactory) NewTable() Table {
	return &modernIkeaTable{
		details: TableDetails{
			Manufacturer: "ikea",
			Type:         "modern",
			LegCount:     4,
			Material:     "wood",
			Color:        "beige",
		},
		price: 249.99,
	}
}

type classicIkeaChair struct {
	details ChairDetails
	price   float64
}

func (c classicIkeaChair) Details() ChairDetails { return c.details }

func (c classicIkeaChair) Price() float64 { return c.price }

type classicIkeaTable struct {
	details TableDetails
	price   float64
}

func (c classicIkeaTable) Details() TableDetails { return c.details }

func (c classicIkeaTable) Price() float64 { return c.price }

type modernIkeaChair struct {
	details ChairDetails
	price   float64
}

func (m modernIkeaChair) Details() ChairDetails { return m.details }

func (m modernIkeaChair) Price() float64 { return m.price }

type modernIkeaTable struct {
	details TableDetails
	price   float64
}

func (m modernIkeaTable) Details() TableDetails { return m.details }

func (m modernIkeaTable) Price() float64 { return m.price }
