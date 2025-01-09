package throw

// ErrorType represents the default type of an error.
type ErrorType string

const (
	ConflictErrorType     ErrorType = "Conflict"
	ForbiddenErrorType    ErrorType = "Forbidden"
	InternalErrorType     ErrorType = "Internal"
	NotFoundErrorType     ErrorType = "Not Found"
	UnauthorizedErrorType ErrorType = "Unauthorized"
	ValidationErrorType   ErrorType = "Validation"
)

func (errorType ErrorType) String() string {
	return string(errorType)
}

// Conflict returns a new [ErrorBuilder] with the Conflict error type.
func Conflict() *ErrorBuilder {
	return NewErrorBuilder(ConflictErrorType.String())
}

// Forbidden returns a new [ErrorBuilder] with the Forbidden error type.
func Forbidden() *ErrorBuilder {
	return NewErrorBuilder(ForbiddenErrorType.String())
}

// Internal returns a new [ErrorBuilder] with the Internal error type.
func Internal() *ErrorBuilder {
	return NewErrorBuilder(InternalErrorType.String())
}

// NotFound returns a new [ErrorBuilder] with the Not Found error type.
func NotFound() *ErrorBuilder {
	return NewErrorBuilder(NotFoundErrorType.String())
}

// Unauthorized returns a new [ErrorBuilder] with the Unauthorized error type.
func Unauthorized() *ErrorBuilder {
	return NewErrorBuilder(UnauthorizedErrorType.String())
}

// Validation returns a new [ErrorBuilder] with the Validation error type.
func Validation() *ErrorBuilder {
	return NewErrorBuilder(ValidationErrorType.String())
}
