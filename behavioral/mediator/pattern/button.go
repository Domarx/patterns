package pattern

func NewButton(mediator Mediator) *Button {
	return &Button{dialog: mediator}
}

type Button struct {
	dialog Mediator
}

func (b *Button) Click() {
	b.dialog.Notify(b, Click)
}

func (b *Button) KeyPress() {
	b.dialog.Notify(b, KeyPress)
}
