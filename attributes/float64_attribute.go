package attributes

import (
	"fmt"
)

// Float64 represents an attribute with float64 type. It will be converted to a string using `%.2f`.
type Float64 struct {
	key   string
	value float64
}

func NewFloat64(key string, value float64) *Float64 {
	return &Float64{key, value}
}

func (attribute *Float64) Key() string {
	return attribute.key
}

func (attribute *Float64) Value() string {
	return fmt.Sprintf("%.2f", attribute.value)
}
