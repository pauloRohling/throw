package throw

import (
	"errors"
)

type TestingT interface {
	Errorf(format string, args ...interface{})
}

// AssertType asserts that the error is of the expected type.
//
// If the error is not of the expected type, it fails the test.
// If the error is nil, it returns true.
func AssertType(t TestingT, err error, expectedType string) {
	if err == nil {
		return
	}

	var customError *Error
	if !errors.As(err, &customError) {
		t.Errorf("expected: %q\n"+"got: %s", expectedType, err.Error())
		return
	}

	if customError.Type() != expectedType {
		t.Errorf("expected: %q\n"+"got: %s", expectedType, customError.Type())
	}
}
