package main

import (
	"encoding/json"
	"os"

	"github.com/domarx/patterns/creational/factory/products"
)

type ProductList struct {
	products []products.Product
}

func (p *ProductList) Add(product products.Product) { p.products = append(p.products, product) }

func (p ProductList) GetCheck() Check {
	check := Check{
		Items: make([]struct {
			Type string  `json:"type"`
			Cost float64 `json:"cost"`
		}, len(p.products)),
		Total: 0,
	}
	for i, product := range p.products {
		check.Items[i].Type = product.Type()
		check.Items[i].Cost = product.Cost()
		check.Total += product.Cost()
	}
	return check
}

type Check struct {
	Items []struct {
		Type string  `json:"type"`
		Cost float64 `json:"cost"`
	} `json:"items"`
	Total float64 `json:"total"`
}

func PrintCheck(check Check) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(check)
}

func main() {
	milkFactory := products.NewMilkFactory()
	breadFactory := products.NewBreadFactory()

	productList := ProductList{}
	productList.Add(milkFactory.NewProduct())
	productList.Add(breadFactory.NewProduct())

	PrintCheck(productList.GetCheck())
}
