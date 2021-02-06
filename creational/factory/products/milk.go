package products

func NewMilkFactory() Factory {
	return &milkFactory{}
}

type milkFactory struct{}

func (m milkFactory) NewProduct() Product { return &milk{} }

type milk struct{}

func (m milk) Type() string { return "milk" }

func (m milk) Cost() float64 { return 1.23 }
