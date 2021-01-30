package factory

type vehicleType int

const (
	Car vehicleType = iota
	Boat
)

type Vehicle interface {
	Drive() error
	Stop() error
}

func CreateVehicle(kind vehicleType) (v Vehicle, err error) {
	switch kind {
	case Car:
		v = &car{}
	case Boat:
		v = &boat{}
	}
	return v, err
}
