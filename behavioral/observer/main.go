package main

import (
	"os"
	"sync"

	"github.com/Sirupsen/logrus"
)

type Event int

const (
	_ Event = iota
	Open
	Write
)

type EventManager struct {
	mux       *sync.Mutex
	listeners map[Event][]EventListener
}

func (e *EventManager) Add(event Event, listener EventListener) {
	e.mux.Lock()
	defer e.mux.Unlock()

	e.listeners[event] = append(e.listeners[event], listener)
}

func (e *EventManager) Remove(event Event, listener EventListener) {
	e.mux.Lock()
	defer e.mux.Unlock()

	for i := 0; i < len(e.listeners[event]); i++ {
		if e.listeners[event][i] == listener {
			e.listeners[event] = append(e.listeners[event][:i], e.listeners[event][i+1:]...)
		}
	}
}

func (e *EventManager) Notify(event Event, path string) {
	for _, listener := range e.listeners[event] {
		listener.Notify(path)
	}
}

func NewEditor() *Editor {
	return &Editor{
		manager: &EventManager{
			mux:       new(sync.Mutex),
			listeners: make(map[Event][]EventListener),
		},
	}
}

type Editor struct {
	manager *EventManager
	file    *os.File
}

func (e *Editor) OpenFile(path string) (err error) {
	e.file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err == nil {
		e.manager.Notify(Open, path)
	}
	return err
}

func (e *Editor) SaveFile() error {
	_, err := e.file.WriteString("test")
	if err == nil {
		stat, statErr := e.file.Stat()
		if statErr != nil {
			return statErr
		}
		e.manager.Notify(Write, stat.Name())
	}
	return err
}

type EventListener interface {
	Notify(path string)
}

func NewLogListener(message string) EventListener {
	return &LoggerListener{message: message}
}

type LoggerListener struct {
	message string
}

func (l *LoggerListener) Notify(path string) {
	logrus.WithField("path", path).Infof(l.message)
}

func main() {
	editor := NewEditor()
	rLog := NewLogListener("file has been opened")
	editor.manager.Add(Open, rLog)

	wLog := NewLogListener("file has been written")
	editor.manager.Add(Write, wLog)

	_ = editor.OpenFile("hello.txt")
	_ = editor.SaveFile()
	editor.manager.Remove(Open, rLog)
	editor.manager.Remove(Write, wLog)
}
