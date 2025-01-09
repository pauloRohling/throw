package attributes

import (
	"encoding/json"
)

type Json struct {
	key   string
	value any
}

func NewJson(key string, value any) *Json {
	return &Json{key, value}
}

func (attribute *Json) Key() string {
	return attribute.key
}

func (attribute *Json) Value() string {
	value, err := json.Marshal(attribute.value)
	if err != nil {
		return "INVALID_JSON_VALUE"
	}
	return string(value)
}
