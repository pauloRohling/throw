package throw

import (
	"errors"
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

// Err sets the error and returns the builder for chaining. If the error is a [throw.Error], it will
// add all its attributes to the current error.
func (b *ErrorBuilder) Err(err error) *ErrorBuilder {
	if err == nil {
		return b
	}

	b.err.err = err
	b.err.message = err.Error()

	var throwError *Error
	if !errors.As(err, &throwError) {
		return b
	}

	for key, attr := range throwError.attributes {
		if b.err.attributes[key] == nil {
			b.err.attributes[key] = attr
		}
	}

	return b
}

// PlainErr sets the error and returns the builder for chaining.
func (b *ErrorBuilder) PlainErr(err error) *ErrorBuilder {
	if err != nil {
		b.err.err = err
		b.err.message = err.Error()
	}

	return b
}

// Str adds a [attributes.String] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Str(key, value string) *ErrorBuilder {
	attr := attributes.NewString(key, value)
	b.err.add(attr)
	return b
}

// Any adds a [attributes.Any] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Any(key string, value any) *ErrorBuilder {
	attr := attributes.NewAny(key, value)
	b.err.add(attr)
	return b
}

// Json adds a [attributes.Json] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Json(key string, value any) *ErrorBuilder {
	attr := attributes.NewJson(key, value)
	b.err.add(attr)
	return b
}

// Uint adds a [attributes.Uint] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Uint(key string, value uint) *ErrorBuilder {
	attr := attributes.NewUint(key, value)
	b.err.add(attr)
	return b
}

// Uint8 adds a [attributes.Uint8] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Uint8(key string, value uint8) *ErrorBuilder {
	attr := attributes.NewUint8(key, value)
	b.err.add(attr)
	return b
}

// Uint16 adds a [attributes.Uint16] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Uint16(key string, value uint16) *ErrorBuilder {
	attr := attributes.NewUint16(key, value)
	b.err.add(attr)
	return b
}

// Uint32 adds a [attributes.Uint32] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Uint32(key string, value uint32) *ErrorBuilder {
	attr := attributes.NewUint32(key, value)
	b.err.add(attr)
	return b
}

// Uint64 adds a [attributes.Uint64] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Uint64(key string, value uint64) *ErrorBuilder {
	attr := attributes.NewUint64(key, value)
	b.err.add(attr)
	return b
}

// Int adds a [attributes.Int] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Int(key string, value int) *ErrorBuilder {
	attr := attributes.NewInt(key, value)
	b.err.add(attr)
	return b
}

// Int8 adds a [attributes.Int8] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Int8(key string, value int8) *ErrorBuilder {
	attr := attributes.NewInt8(key, value)
	b.err.add(attr)
	return b
}

// Int16 adds a [attributes.Int16] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Int16(key string, value int16) *ErrorBuilder {
	attr := attributes.NewInt16(key, value)
	b.err.add(attr)
	return b
}

// Int32 adds a [attributes.Int32] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Int32(key string, value int32) *ErrorBuilder {
	attr := attributes.NewInt32(key, value)
	b.err.add(attr)
	return b
}

// Int64 adds a [attributes.Int64] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Int64(key string, value int64) *ErrorBuilder {
	attr := attributes.NewInt64(key, value)
	b.err.add(attr)
	return b
}

// Float32 adds a [attributes.Float32] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Float32(key string, value float32) *ErrorBuilder {
	attr := attributes.NewFloat32(key, value)
	b.err.add(attr)
	return b
}

// Float64 adds a [attributes.Float64] attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Float64(key string, value float64) *ErrorBuilder {
	attr := attributes.NewFloat64(key, value)
	b.err.add(attr)
	return b
}

// Attribute adds an attribute to the error and returns the builder for chaining.
func (b *ErrorBuilder) Attribute(attribute Attribute) *ErrorBuilder {
	b.err.add(attribute)
	return b
}

// Attr adds an attribute to the error and returns the builder for chaining. It is a shorthand for [ErrorBuilder.Attribute].
func (b *ErrorBuilder) Attr(attribute Attribute) *ErrorBuilder {
	b.err.add(attribute)
	return b
}

// Msg sets the error message and returns the error. It ends the chain of methods.
func (b *ErrorBuilder) Msg(message string) *Error {
	b.err.message = message
	return b.err
}

// Msgf sets the error message with a formatted string and returns the error. It ends the chain of methods.
func (b *ErrorBuilder) Msgf(message string, values ...any) *Error {
	b.err.message = fmt.Sprintf(message, values...)
	return b.err
}
