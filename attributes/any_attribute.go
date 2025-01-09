package attributes

import (
	"fmt"
)

// Any represents an attribute with any type. It will be converted to a string using `%+v`.
type Any struct {
	key   string
	value any
}

func NewAny(key string, value any) *Any {
	return &Any{key, value}
}

func (attribute *Any) Key() string {
	return attribute.key
}

func (attribute *Any) Value() string {
	return fmt.Sprintf("%+v", attribute.value)
}
