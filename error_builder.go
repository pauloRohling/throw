package throw

import (
	"fmt"
	"github.com/pauloRohling/throw/attributes"
)

// ErrorBuilder is responsible for building [Error] objects.
//
// It has methods for adding attributes to the error, which can be used to provide more context
// about the error.
//
// It also has methods for setting the error message and returning the built error, which should
// be used at the end of the chain of methods.
type ErrorBuilder struct {
	err *Error
}

// NewErrorBuilder creates a new [ErrorBuilder] without error type.
func NewErrorBuilder() *ErrorBuilder {
	return &ErrorBuilder{
		err: newError(""),
	}
}

// NewTypedErrorBuilder creates a new [ErrorBuilder] with the given error type.
func NewTypedErrorBuilder(errType string) *ErrorBuilder {
	return &ErrorBuilder{
		err: newError(errType),
	}
}

// Err sets the error and returns the builder for chaining.
func (builder *ErrorBuilder) Err(err error) *ErrorBuilder {
	if err != nil {
		builder.err.err = err
		builder.err.message = err.Error()
	}

	return builder
}

// Str adds a [attributes.String] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Str(key, value string) *ErrorBuilder {
	attr := attributes.NewString(key, value)
	builder.err.add(attr)
	return builder
}

// Any adds a [attributes.Any] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Any(key string, value any) *ErrorBuilder {
	attr := attributes.NewAny(key, value)
	builder.err.add(attr)
	return builder
}

// Json adds a [attributes.Json] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Json(key string, value any) *ErrorBuilder {
	attr := attributes.NewJson(key, value)
	builder.err.add(attr)
	return builder
}

// Uint adds a [attributes.Uint] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Uint(key string, value uint) *ErrorBuilder {
	attr := attributes.NewUint(key, value)
	builder.err.add(attr)
	return builder
}

// Uint8 adds a [attributes.Uint8] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Uint8(key string, value uint8) *ErrorBuilder {
	attr := attributes.NewUint8(key, value)
	builder.err.add(attr)
	return builder
}

// Uint16 adds a [attributes.Uint16] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Uint16(key string, value uint16) *ErrorBuilder {
	attr := attributes.NewUint16(key, value)
	builder.err.add(attr)
	return builder
}

// Uint32 adds a [attributes.Uint32] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Uint32(key string, value uint32) *ErrorBuilder {
	attr := attributes.NewUint32(key, value)
	builder.err.add(attr)
	return builder
}

// Uint64 adds a [attributes.Uint64] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Uint64(key string, value uint64) *ErrorBuilder {
	attr := attributes.NewUint64(key, value)
	builder.err.add(attr)
	return builder
}

// Int adds a [attributes.Int] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Int(key string, value int) *ErrorBuilder {
	attr := attributes.NewInt(key, value)
	builder.err.add(attr)
	return builder
}

// Int8 adds a [attributes.Int8] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Int8(key string, value int8) *ErrorBuilder {
	attr := attributes.NewInt8(key, value)
	builder.err.add(attr)
	return builder
}

// Int16 adds a [attributes.Int16] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Int16(key string, value int16) *ErrorBuilder {
	attr := attributes.NewInt16(key, value)
	builder.err.add(attr)
	return builder
}

// Int32 adds a [attributes.Int32] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Int32(key string, value int32) *ErrorBuilder {
	attr := attributes.NewInt32(key, value)
	builder.err.add(attr)
	return builder
}

// Int64 adds a [attributes.Int64] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Int64(key string, value int64) *ErrorBuilder {
	attr := attributes.NewInt64(key, value)
	builder.err.add(attr)
	return builder
}

// Float32 adds a [attributes.Float32] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Float32(key string, value float32) *ErrorBuilder {
	attr := attributes.NewFloat32(key, value)
	builder.err.add(attr)
	return builder
}

// Float64 adds a [attributes.Float64] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Float64(key string, value float64) *ErrorBuilder {
	attr := attributes.NewFloat64(key, value)
	builder.err.add(attr)
	return builder
}

// Msg sets the error message and returns the error. It ends the chain of methods.
func (builder *ErrorBuilder) Msg(message string) *Error {
	builder.err.message = message
	return builder.err
}

// Msgf sets the error message with a formatted string and returns the error. It ends the chain of methods.
func (builder *ErrorBuilder) Msgf(message string, values ...any) *Error {
	builder.err.message = fmt.Sprintf(message, values...)
	return builder.err
}
