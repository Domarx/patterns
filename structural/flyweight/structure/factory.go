package structure

import "sync"

var (
	mux   = new(sync.Mutex)
	store = make([]*flyweight, 0)
)

// getFlyweight returns a flyweight object from the cache if found, otherwise
// new flyweight object is created. This way we store as few flyweight references
// as possible.
func getFlyweight(name, color, texture string) *flyweight {
	mux.Lock()
	defer mux.Unlock()

	var f *flyweight
	for _, obj := range store {
		if obj.name == name && obj.color == color && obj.texture == texture {
			f = obj
		}
	}
	if f == nil {
		f = newFlyweight(name, color, texture)
		f.id = len(store) + 1
		store = append(store, f)
	}
	return f
}
