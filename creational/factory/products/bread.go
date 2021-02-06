package products

func NewBreadFactory() Factory {
	return &breadFactory{}
}

type breadFactory struct{}

func (b breadFactory) NewProduct() Product { return &bread{} }

type bread struct{}

func (b bread) Type() string { return "bread" }

func (b bread) Cost() float64 { return 0.99 }
