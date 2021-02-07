package furniture

type jyskFactory struct{}

func (j jyskFactory) GetFactory(factory string) Factory {
	var f Factory
	switch factory {
	case "outdoor":
		f = &outdoorJyskFactory{}
	case "indoor":
		f = &indoorJyskFactory{}
	}
	return f
}

type outdoorJyskFactory struct{}

func (o outdoorJyskFactory) NewChair() Chair {
	return &outdoorJyskChair{
		details: ChairDetails{
			Manufacturer: "jysk",
			Type:         "outdoor",
			LegCount:     4,
			Material:     "plastic",
			Color:        "white",
		},
		price: 25.00,
	}
}

func (o outdoorJyskFactory) NewTable() Table {
	return &outdoorJyskTable{
		details: TableDetails{
			Manufacturer: "jysk",
			Type:         "outdoor",
			LegCount:     4,
			Material:     "plastic",
			Color:        "white",
		},
		price: 75.00,
	}
}

type outdoorJyskChair struct {
	details ChairDetails
	price   float64
}

func (o outdoorJyskChair) Details() ChairDetails { return o.details }

func (o outdoorJyskChair) Price() float64 { return o.price }

type outdoorJyskTable struct {
	details TableDetails
	price   float64
}

func (o outdoorJyskTable) Details() TableDetails { return o.details }

func (o outdoorJyskTable) Price() float64 { return o.price }

type indoorJyskFactory struct{}

func (i indoorJyskFactory) NewChair() Chair {
	return &indoorJyskChair{
		details: ChairDetails{
			Manufacturer: "jysk",
			Type:         "indoor",
			LegCount:     3,
			Material:     "leather",
			Color:        "blue",
		},
		price: 50.00,
	}
}

func (i indoorJyskFactory) NewTable() Table {
	return &indoorJyskTable{
		details: TableDetails{
			Manufacturer: "jysk",
			Type:         "indoor",
			LegCount:     4,
			Material:     "leather",
			Color:        "red",
		},
		price: 199.00,
	}
}

type indoorJyskChair struct {
	details ChairDetails
	price   float64
}

func (i indoorJyskChair) Details() ChairDetails { return i.details }

func (i indoorJyskChair) Price() float64 { return i.price }

type indoorJyskTable struct {
	details TableDetails
	price   float64
}

func (i indoorJyskTable) Details() TableDetails { return i.details }

func (i indoorJyskTable) Price() float64 { return i.price }
