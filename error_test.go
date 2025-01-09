package throw

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

	err := NewTypedErrorBuilder("testErrorType").
		Err(someError).
		Str("key", "value").
		Any("struct", aStructValue).
		Json("json", aStructValue).
		Int("age", 20).
		Int8("age8", 8).
		Int16("age16", 16).
		Int32("age32", 32).
		Int64("age64", 64).
		Uint("uage", 20).
		Uint8("uage8", 8).
		Uint16("uage16", 16).
		Uint32("uage32", 32).
		Uint64("uage64", 64).
		Float32("float32", 20.32).
		Float64("float64", 20.32).
		Msg("Some error message")

	errFormatted := NewErrorBuilder().
		Msgf("Some error message %s", "value")

	assert.Equal(t, "Some error message", err.Error())
	assert.Equal(t, "Some error message value", errFormatted.Error())
	assert.Equal(t, "testErrorType", err.Type())
	assert.Equal(t, someError, err.Unwrap())

	attrs := err.Attributes()
	assert.Equal(t, 15, len(attrs))
	assert.IsType(t, &attributes.String{}, attrs[0])
	assert.IsType(t, &attributes.Any{}, attrs[1])
	assert.IsType(t, &attributes.Json{}, attrs[2])
	assert.IsType(t, &attributes.Int{}, attrs[3])
	assert.IsType(t, &attributes.Int8{}, attrs[4])
	assert.IsType(t, &attributes.Int16{}, attrs[5])
	assert.IsType(t, &attributes.Int32{}, attrs[6])
	assert.IsType(t, &attributes.Int64{}, attrs[7])
	assert.IsType(t, &attributes.Uint{}, attrs[8])
	assert.IsType(t, &attributes.Uint8{}, attrs[9])
	assert.IsType(t, &attributes.Uint16{}, attrs[10])
	assert.IsType(t, &attributes.Uint32{}, attrs[11])
	assert.IsType(t, &attributes.Uint64{}, attrs[12])
	assert.IsType(t, &attributes.Float32{}, attrs[13])
	assert.IsType(t, &attributes.Float64{}, attrs[14])
}
