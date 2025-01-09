package main

import (
	"errors"
	"github.com/pauloRohling/throw/attributes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type AStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestError(t *testing.T) {
	aStructValue := AStruct{Name: "Lorem Ipsum", Age: 20}

	someError := errors.New("some error")

	err := NewErrorBuilder("testErrorType").
		Err(someError).
		Str("key", "value").
		Any("struct", aStructValue).
		Json("json", aStructValue).
		Int("age", 20).
		Msg("Some error message")

	errFormatted := NewErrorBuilder("testErrorType").
		Msgf("Some error message %s", "value")

	assert.Equal(t, "Some error message", err.Error())
	assert.Equal(t, "Some error message value", errFormatted.Error())
	assert.Equal(t, "testErrorType", err.Type())
	assert.Equal(t, someError, err.Unwrap())

	attrs := err.Attributes()
	assert.Equal(t, 4, len(attrs))
	assert.IsType(t, &attributes.String{}, attrs[0])
	assert.IsType(t, &attributes.Any{}, attrs[1])
	assert.IsType(t, &attributes.Json{}, attrs[2])
	assert.IsType(t, &attributes.Int{}, attrs[3])
}
