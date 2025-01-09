package attributes

import (
	"fmt"
)

// Int32 represents an attribute with int32 type.
type Int32 struct {
	key   string
	value int32
}

func NewInt32(key string, value int32) *Int32 {
	return &Int32{key, value}
}

func (attribute *Int32) Key() string {
	return attribute.key
}

func (attribute *Int32) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
