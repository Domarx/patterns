package structure

import (
	"encoding/json"
	"fmt"
)

// flyweight should contain intrinsic (immutable) values which are
// read only. Flyweight state stays within the object, letting us reuse
// it in different contexts. This approach allows us to use fewer
// flyweight objects, which have fewer variations than the objects
// with extrinsic values. Mutable values are no longer stored inside the
// flyweight, but rather they are passed to the object method that relies on them.
type flyweight struct {
	id      int // id is only to check if objects are being reused
	name    string
	color   string
	texture string
}

// Draw takes mutable values and performs an action
func (f *flyweight) Draw(x, y int) {
	raw, _ := json.Marshal(map[string]interface{}{
		"context": map[string]int{
			"x": x,
			"y": y,
		},
		"flyweight": map[string]interface{}{
			"id":      f.id,
			"name":    f.name,
			"color":   f.color,
			"texture": f.texture,
		},
	})
	fmt.Println(string(raw))
}

func newFlyweight(name, color, texture string) *flyweight {
	return &flyweight{
		name:    name,
		color:   color,
		texture: texture,
	}
}
