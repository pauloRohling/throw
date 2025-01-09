package main

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

// NewErrorBuilder creates a new [ErrorBuilder] with the given error type.
func NewErrorBuilder(errType string) *ErrorBuilder {
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

// Int adds a [attributes.Int] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Int(key string, value int) *ErrorBuilder {
	attr := attributes.NewInt(key, value)
	builder.err.add(attr)
	return builder
}

// Json adds a [attributes.Json] attribute to the error and returns the builder for chaining.
func (builder *ErrorBuilder) Json(key string, value any) *ErrorBuilder {
	attr := attributes.NewJson(key, value)
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
