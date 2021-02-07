package main

import (
	"encoding/json"
	"os"

	"github.com/domarx/patterns/creational/builder/house"
)

func main() {
	contractor := house.Contractor{}

	h, _ := contractor.BuildHouse(house.Simple)
	PrintHouse(h)

	h, _ = contractor.BuildHouse(house.Rich)
	PrintHouse(h)
}

func PrintHouse(house house.House) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(house)
}
