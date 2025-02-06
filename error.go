package throw

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
func (error *Error) Error() string {
	return error.message
}

// Unwrap returns the underlying error.
func (error *Error) Unwrap() error {
	return error.err
}

// Attributes returns the attributes of the error.
func (error *Error) Attributes() []Attribute {
	attributes := make([]Attribute, 0, len(error.attributes))
	for _, attribute := range error.attributes {
		attributes = append(attributes, attribute)
	}
	return attributes
}

// Attribute returns the attribute with the given key.
// If the attribute does not exist, nil is returned.
func (error *Error) Attribute(key string) Attribute {
	return error.attributes[key]
}

// Type returns the type of the error.
func (error *Error) Type() string {
	return error.errType
}

func (error *Error) add(attribute Attribute) {
	error.attributes[attribute.Key()] = attribute
}
