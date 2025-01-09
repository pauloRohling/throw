package attributes

import (
	"fmt"
)

// Int64 represents an attribute with int64 type.
type Int64 struct {
	key   string
	value int64
}

func NewInt64(key string, value int64) *Int64 {
	return &Int64{key, value}
}

func (attribute *Int64) Key() string {
	return attribute.key
}

func (attribute *Int64) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
