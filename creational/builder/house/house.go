package house

type HouseType int

const (
	Simple HouseType = iota
	Rich
)

type House struct {
	Floor    string
	Walls    string
	Roof     string
	Windows  string
	Pool     bool
	BBQGrill bool
}
