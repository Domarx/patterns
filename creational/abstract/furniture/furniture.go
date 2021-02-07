package furniture

type ProducerType int

const (
	Ikea ProducerType = iota
	Jysk
)

func GetFurnitureFactory(producerType ProducerType) Producer {
	var p Producer
	switch producerType {
	case Ikea:
		p = &ikeaFactory{}
	case Jysk:
		p = &jyskFactory{}
	}
	return p
}

type Producer interface {
	GetFactory(factory string) Factory
}

type Factory interface {
	NewChair() Chair
	NewTable() Table
}

type Chair interface {
	Details() ChairDetails
	Price() float64
}

type ChairDetails struct {
	Manufacturer string
	Type         string
	LegCount     int
	Material     string
	Color        string
}

type Table interface {
	Details() TableDetails
	Price() float64
}

type TableDetails struct {
	Manufacturer string
	Type         string
	LegCount     int
	Material     string
	Color        string
}
