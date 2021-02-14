package device

import (
	"fmt"
)

type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	Volume() float64
	SetVolume(percent float64)
}

func NewTV() Device { return &tv{} }

type tv struct {
	enabled bool
	volume  float64
}

func (t *tv) IsEnabled() bool { return t.enabled }

func (t *tv) Enable() {
	fmt.Println("TV is ON")
	t.enabled = true
}

func (t *tv) Disable() {
	fmt.Println("TV is OFF")
	t.enabled = false
}

func (t *tv) Volume() float64 { return t.volume }

func (t *tv) SetVolume(percent float64) {
	t.volume = percent
	fmt.Println("setting TV volume to: ", t.volume)
}

func NewRadio() Device { return &radio{} }

type radio struct {
	enabled bool
	volume  float64
}

func (r *radio) IsEnabled() bool { return r.enabled }

func (r *radio) Enable() {
	fmt.Println("Radio is ON")
	r.enabled = true
}

func (r *radio) Disable() {
	fmt.Println("Radio is OFF")
	r.enabled = false
}

func (r *radio) Volume() float64 { return r.volume }

func (r *radio) SetVolume(percent float64) {
	r.volume = percent
	fmt.Println("setting Radio volume to: ", r.volume)
}
