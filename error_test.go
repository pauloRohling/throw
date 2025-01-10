package throw

import (
	"errors"
	"github.com/pauloRohling/throw/attributes"
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

	assertEqual(t, "Some error message", err.Error())
	assertEqual(t, "Some error message value", errFormatted.Error())
	assertEqual(t, "testErrorType", err.Type())
	AssertType(t, err, "testErrorType")
	assertEqual(t, someError, err.Unwrap())

	attrs := err.Attributes()
	assertEqual(t, 15, len(attrs))
	assertIsType(t, &attributes.String{}, attrs[0])
	assertIsType(t, &attributes.Any{}, attrs[1])
	assertIsType(t, &attributes.Json{}, attrs[2])
	assertIsType(t, &attributes.Int{}, attrs[3])
	assertIsType(t, &attributes.Int8{}, attrs[4])
	assertIsType(t, &attributes.Int16{}, attrs[5])
	assertIsType(t, &attributes.Int32{}, attrs[6])
	assertIsType(t, &attributes.Int64{}, attrs[7])
	assertIsType(t, &attributes.Uint{}, attrs[8])
	assertIsType(t, &attributes.Uint8{}, attrs[9])
	assertIsType(t, &attributes.Uint16{}, attrs[10])
	assertIsType(t, &attributes.Uint32{}, attrs[11])
	assertIsType(t, &attributes.Uint64{}, attrs[12])
	assertIsType(t, &attributes.Float32{}, attrs[13])
	assertIsType(t, &attributes.Float64{}, attrs[14])
}
