package house

type Contractor struct{}

func (c Contractor) BuildHouse(houseType HouseType) (h House, err error) {
	b := &baseBuilder{}
	switch houseType {
	case Simple:
		h, err = simpleHouseBuilder{builder: b}.buildHouse()
	case Rich:
		h, err = richHouseBuilder{builder: b}.buildHouse()
	}
	return h, err
}
