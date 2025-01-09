package attributes

import (
	"fmt"
)

// Float32 represents an attribute with float32 type. It will be converted to a string using `%.2f`.
type Float32 struct {
	key   string
	value float32
}

func NewFloat32(key string, value float32) *Float32 {
	return &Float32{key, value}
}

func (attribute *Float32) Key() string {
	return attribute.key
}

func (attribute *Float32) Value() string {
	return fmt.Sprintf("%.2f", attribute.value)
}
