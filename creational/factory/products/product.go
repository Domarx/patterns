package products

type Factory interface {
	NewProduct() Product
}

type Product interface {
	Type() string
	Cost() float64
}
