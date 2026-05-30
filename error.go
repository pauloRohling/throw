package throw

import (
	"errors"
	"iter"
)

// Attribute represents an attribute that can be added to an error.
type Attribute interface {
	Key() string
	Value() string
}

// Error represents an error that can be built using the [ErrorBuilder].
type Error struct {
	attributes map[string]Attribute
	message    string
	errType    string
	err        error
}

func newError(errType string) *Error {
	return &Error{
		errType:    errType,
		attributes: make(map[string]Attribute),
	}
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.message
}

// Unwrap returns the underlying error.
func (e *Error) Unwrap() error {
	return e.err
}

// UnwrapOriginal returns the original error.
func (e *Error) UnwrapOriginal() error {
	originalError := e.err

	for originalError != nil {
		var unwrappedError *Error
		if !errors.As(originalError, &unwrappedError) {
			break
		}
		originalError = unwrappedError.err
	}

	return originalError
}

// Attributes returns an iterator that yields the error attributes.
func (e *Error) Attributes() iter.Seq[Attribute] {
	return func(yield func(Attribute) bool) {
		for _, value := range e.attributes {
			if !yield(value) {
				return
			}
		}
	}
}

// Attribute returns the attribute with the given key.
// If the attribute does not exist, nil is returned.
func (e *Error) Attribute(key string) Attribute {
	return e.attributes[key]
}

// Type returns the type of the error.
func (e *Error) Type() string {
	return e.errType
}

func (e *Error) add(attribute Attribute) {
	e.attributes[attribute.Key()] = attribute
}
