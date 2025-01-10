package throw

import "net/http"

// ErrorType is an example of how an error type can be defined.
// It has types that can be mapped to HTTP status codes.
type ErrorType string

const (
	ConflictErrorType     ErrorType = "Conflict"     // ConflictErrorType is an error type that represents a conflict (409)
	ForbiddenErrorType    ErrorType = "Forbidden"    // ForbiddenErrorType is an error type that represents a forbidden access (403)
	NotFoundErrorType     ErrorType = "Not Found"    // NotFoundErrorType is an error type that represents a not found error (404)
	UnauthorizedErrorType ErrorType = "Unauthorized" // UnauthorizedErrorType is an error type that represents an unauthorized access (401)
	ValidationErrorType   ErrorType = "Validation"   // ValidationErrorType is an error type that represents a validation error (400)
	InternalErrorType     ErrorType = "Internal"     // InternalErrorType is an error type that represents an internal server error (500)
)

// String returns the string representation of the ErrorType
func (errorType ErrorType) String() string {
	return string(errorType)
}

// StatusCode returns the HTTP status code associated with the ErrorType.
// Defaults to http.StatusInternalServerError
func (errorType ErrorType) StatusCode() int {
	switch errorType {
	case ConflictErrorType:
		return http.StatusConflict
	case ForbiddenErrorType:
		return http.StatusForbidden
	case NotFoundErrorType:
		return http.StatusNotFound
	case UnauthorizedErrorType:
		return http.StatusUnauthorized
	case ValidationErrorType:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

// Conflict returns a new [ErrorBuilder] with the Conflict error type.
func Conflict() *ErrorBuilder {
	return NewTypedErrorBuilder(ConflictErrorType.String())
}

// Forbidden returns a new [ErrorBuilder] with the Forbidden error type.
func Forbidden() *ErrorBuilder {
	return NewTypedErrorBuilder(ForbiddenErrorType.String())
}

// Internal returns a new [ErrorBuilder] with the Internal error type.
func Internal() *ErrorBuilder {
	return NewTypedErrorBuilder(InternalErrorType.String())
}

// NotFound returns a new [ErrorBuilder] with the Not Found error type.
func NotFound() *ErrorBuilder {
	return NewTypedErrorBuilder(NotFoundErrorType.String())
}

// Unauthorized returns a new [ErrorBuilder] with the Unauthorized error type.
func Unauthorized() *ErrorBuilder {
	return NewTypedErrorBuilder(UnauthorizedErrorType.String())
}

// Validation returns a new [ErrorBuilder] with the Validation error type.
func Validation() *ErrorBuilder {
	return NewTypedErrorBuilder(ValidationErrorType.String())
}
