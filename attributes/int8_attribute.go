package attributes

import (
	"fmt"
)

// Int8 represents an attribute with int8 type.
type Int8 struct {
	key   string
	value int8
}

func NewInt8(key string, value int8) *Int8 {
	return &Int8{key, value}
}

func (attribute *Int8) Key() string {
	return attribute.key
}

func (attribute *Int8) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
