package attributes

import (
	"fmt"
)

// Int16 represents an attribute with int16 type.
type Int16 struct {
	key   string
	value int16
}

func NewInt16(key string, value int16) *Int16 {
	return &Int16{key, value}
}

func (attribute *Int16) Key() string {
	return attribute.key
}

func (attribute *Int16) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
