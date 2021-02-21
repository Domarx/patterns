package main

import "fmt"

type Snapshot interface {
	Restore()
}

func NewSnapshot(editor *Editor, state string) Snapshot {
	return &snapshot{
		editor: editor,
		state:  state,
	}
}

type snapshot struct {
	editor *Editor
	state  string
}

func (s *snapshot) Restore() {
	s.editor.SetState(s.state)
}

type Editor struct {
	state string
}

func (e *Editor) String() string {
	return fmt.Sprintf("State = [%s]", e.state)
}

func (e *Editor) SetState(state string) { e.state = state }

func (e *Editor) Snapshot() Snapshot {
	return NewSnapshot(e, e.state)
}

func NewCaretaker() *caretaker {
	return &caretaker{}
}

type caretaker struct {
	history []Snapshot
}

func (c *caretaker) Backup(editor *Editor) {
	c.history = append(c.history, editor.Snapshot())
}

func (c *caretaker) Undo() {
	s := c.history[len(c.history)-1]
	c.history = c.history[:len(c.history)-1]
	s.Restore()
}

func NewClient(caretaker *caretaker, editor *Editor) *Client {
	return &Client{
		caretaker: caretaker,
		editor:    editor,
	}
}

type Client struct {
	caretaker *caretaker
	editor    *Editor
}

func (c *Client) Save(state string) {
	c.caretaker.Backup(c.editor)
	c.editor.SetState(state)
}

func (c *Client) Undo() {
	c.caretaker.Undo()
}

func main() {
	editor := &Editor{}
	c := NewClient(NewCaretaker(), editor)
	c.Save("action")
	fmt.Println(editor)

	c.Save("another action")
	fmt.Println(editor)

	c.Undo()
	fmt.Println(editor)
}
