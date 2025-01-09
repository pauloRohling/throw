package attributes

import (
	"fmt"
)

// Int represents an attribute with int type.
type Int struct {
	key   string
	value int
}

func NewInt(key string, value int) *Int {
	return &Int{key, value}
}

func (attribute *Int) Key() string {
	return attribute.key
}

func (attribute *Int) Value() string {
	return fmt.Sprintf("%d", attribute.value)
}
