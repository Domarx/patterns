package main

import "github.com/domarx/patterns/behavioral/mediator/pattern"

func main() {
	var (
		button   = new(pattern.Button)
		checkbox = new(pattern.CheckBox)
	)
	pattern.NewDialog(
		pattern.WithButton(button),
		pattern.WithCheckBox(checkbox),
	)
	{
		button.KeyPress() // component = [*pattern.Button], event = Press
		button.Click()    // component = [*pattern.Button], event = Click
	}
	{
		checkbox.Click()    // component = [*pattern.CheckBox], event = Click
		checkbox.KeyPress() // component = [*pattern.CheckBox], event = Press
		checkbox.Check()    // component = [*pattern.CheckBox], event = Check
	}
}
