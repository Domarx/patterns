package pattern

import "fmt"

type Event int

func (e Event) String() string {
	return [...]string{"Click", "Press", "Check"}[e]
}

const (
	Click Event = iota
	KeyPress
	Check
)

type Component interface{}

type Mediator interface {
	Notify(component Component, event Event)
}

type option func(dialog *dialog)

var (
	WithButton = func(button *Button) option {
		return func(dialog *dialog) {
			button.dialog = dialog
			dialog.button = button
		}
	}
	WithCheckBox = func(box *CheckBox) option {
		return func(dialog *dialog) {
			box.dialog = dialog
			dialog.checkBox = box
		}
	}
)

func NewDialog(options ...option) Mediator {
	a := &dialog{}
	for _, opt := range options {
		opt(a)
	}
	return a
}

type dialog struct {
	button   *Button
	checkBox *CheckBox
}

func (a *dialog) Notify(component Component, event Event) {
	// logic goes here
	fmt.Printf("component = [%T], event = %s\n", component, event)
}
