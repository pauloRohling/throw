package attributes

import (
	"fmt"
)

// Uint64 represents an attribute with uint64 type.
type Uint64 struct {
	key   string
	value uint64
}

func NewUint64(key string, value uint64) *Uint64 {
	return &Uint64{key, value}
}

func (attribute *Uint64) Key() string {
	return attribute.key
}

func (attribute *Uint64) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
