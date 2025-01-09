package main

// Attribute represents an attribute that can be added to an error.
type Attribute interface {
	Key() string
	Value() string
}

// Error represents an error that can be built using the [ErrorBuilder].
type Error struct {
	attributes []Attribute
	message    string
	errType    string
	err        error
}

func newError(errType string) *Error {
	return &Error{errType: errType}
}

// Error returns the error message.
func (error *Error) Error() string {
	return error.message
}

// Unwrap returns the underlying error.
func (error *Error) Unwrap() error {
	return error.err
}

// Attributes returns the attributes of the error.
func (error *Error) Attributes() []Attribute {
	return error.attributes
}

// Type returns the type of the error.
func (error *Error) Type() string {
	return error.errType
}

func (error *Error) add(attribute Attribute) {
	error.attributes = append(error.attributes, attribute)
}
