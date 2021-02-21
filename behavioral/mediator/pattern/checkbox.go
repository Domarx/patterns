package pattern

func NewCheckBox(mediator Mediator) *CheckBox {
	return &CheckBox{dialog: mediator}
}

type CheckBox struct {
	dialog Mediator
}

func (c *CheckBox) Click() {
	c.dialog.Notify(c, Click)
}

func (c *CheckBox) KeyPress() {
	c.dialog.Notify(c, KeyPress)
}

func (c *CheckBox) Check() {
	c.dialog.Notify(c, Check)
}
