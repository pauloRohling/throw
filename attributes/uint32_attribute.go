package attributes

import (
	"fmt"
)

// Uint32 represents an attribute with uint32 type.
type Uint32 struct {
	key   string
	value uint32
}

func NewUint32(key string, value uint32) *Uint32 {
	return &Uint32{key, value}
}

func (attribute *Uint32) Key() string {
	return attribute.key
}

func (attribute *Uint32) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
