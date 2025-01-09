package attributes

import (
	"fmt"
)

// Uint8 represents an attribute with uint8 type.
type Uint8 struct {
	key   string
	value uint8
}

func NewUint8(key string, value uint8) *Uint8 {
	return &Uint8{key, value}
}

func (attribute *Uint8) Key() string {
	return attribute.key
}

func (attribute *Uint8) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
