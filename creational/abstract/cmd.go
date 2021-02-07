package main

import (
	"encoding/json"
	"os"

	"github.com/domarx/patterns/creational/abstract/furniture"
)

type client struct {
	factory furniture.Producer
	chairs  []furniture.Chair
	tables  []furniture.Table
}

func (c *client) AddChair(chairType string) {
	c.chairs = append(c.chairs, c.factory.GetFactory(chairType).NewChair())
}

func (c *client) AddTable(tableType string) {
	c.tables = append(c.tables, c.factory.GetFactory(tableType).NewTable())
}

func (c *client) NewFactory(producerType furniture.ProducerType) {
	c.factory = furniture.GetFurnitureFactory(producerType)
}

func (c *client) GetItems() interface{} {
	chairs := make([]interface{}, len(c.chairs))
	for i, chair := range c.chairs {
		chairs[i] = map[string]interface{}{
			"details": chair.Details(),
			"price":   chair.Price(),
		}
	}
	tables := make([]interface{}, len(c.tables))
	for i, table := range c.tables {
		tables[i] = map[string]interface{}{
			"details": table.Details(),
			"price":   table.Price(),
		}
	}
	items := map[string]interface{}{
		"chairs": chairs,
		"tables": tables,
	}
	return items
}

func PrintItems(items interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(items)
}

func main() {
	c := client{}

	c.NewFactory(furniture.Jysk)
	c.AddChair("indoor")
	c.AddTable("indoor")
	c.AddChair("outdoor")
	c.AddTable("outdoor")

	c.NewFactory(furniture.Ikea)
	c.AddChair("modern")
	c.AddTable("modern")
	c.AddChair("classic")
	c.AddTable("classic")

	PrintItems(c.GetItems())
}
