package attributes

import (
	"fmt"
)

// Uint16 represents an attribute with uint16 type.
type Uint16 struct {
	key   string
	value uint16
}

func NewUint16(key string, value uint16) *Uint16 {
	return &Uint16{key, value}
}

func (attribute *Uint16) Key() string {
	return attribute.key
}

func (attribute *Uint16) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
