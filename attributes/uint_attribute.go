package attributes

import (
	"fmt"
)

// Uint represents an attribute with uint type.
type Uint struct {
	key   string
	value uint
}

func NewUint(key string, value uint) *Uint {
	return &Uint{key, value}
}

func (attribute *Uint) Key() string {
	return attribute.key
}

func (attribute *Uint) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
