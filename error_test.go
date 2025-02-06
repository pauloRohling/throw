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

type AAttribute struct {
	key string
}

func (attribute *AAttribute) Key() string {
	return attribute.key
}

func (attribute *AAttribute) Value() string {
	return "value"
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
		Attribute(&AAttribute{key: "key1"}).
		Attr(&AAttribute{key: "key2"}).
		Msg("Some error message")

	errFormatted := NewErrorBuilder().
		Msgf("Some error message %s", "value")

	assertEqual(t, "Some error message", err.Error())
	assertEqual(t, "Some error message value", errFormatted.Error())
	assertEqual(t, "testErrorType", err.Type())
	AssertType(t, err, "testErrorType")
	assertEqual(t, someError, err.Unwrap())

	attrs := err.Attributes()
	assertEqual(t, 17, len(attrs))

	assertIsType(t, &attributes.String{}, err.Attribute("key"))
	assertIsType(t, &attributes.Any{}, err.Attribute("struct"))
	assertIsType(t, &attributes.Json{}, err.Attribute("json"))
	assertIsType(t, &attributes.Int{}, err.Attribute("age"))
	assertIsType(t, &attributes.Int8{}, err.Attribute("age8"))
	assertIsType(t, &attributes.Int16{}, err.Attribute("age16"))
	assertIsType(t, &attributes.Int32{}, err.Attribute("age32"))
	assertIsType(t, &attributes.Int64{}, err.Attribute("age64"))
	assertIsType(t, &attributes.Uint{}, err.Attribute("uage"))
	assertIsType(t, &attributes.Uint8{}, err.Attribute("uage8"))
	assertIsType(t, &attributes.Uint16{}, err.Attribute("uage16"))
	assertIsType(t, &attributes.Uint32{}, err.Attribute("uage32"))
	assertIsType(t, &attributes.Uint64{}, err.Attribute("uage64"))
	assertIsType(t, &attributes.Float32{}, err.Attribute("float32"))
	assertIsType(t, &attributes.Float64{}, err.Attribute("float64"))
	assertIsType(t, &AAttribute{}, err.Attribute("key1"))
	assertIsType(t, &AAttribute{}, err.Attribute("key2"))
}
