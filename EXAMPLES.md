## Examples

### Creating a simple error

To create a simple error, you can use the `NewErrorBuilder` function, which returns a new `ErrorBuilder` instance
without a type. Throw does not require you to specify a type for your error, but it can be useful.

```go
package fail

import "github.com/pauloRohling/throw"

func AlwaysFail() error {
	return throw.NewErrorBuilder().Msg("An error has occurred!")
}
```

### Creating an error with attributes and custom type

To add attributes to your error, you can use the chainable methods provided by the `ErrorBuilder` instance. It is
possible to get the attributes of the error using the `Attributes()` method.

By using the `NewTypedErrorBuilder` function, you can specify a custom type for your error, which can be retrieved using
the `Type()` method.

```go
package fail

import "github.com/pauloRohling/throw"

func AlwaysFail() error {
	return throw.NewTypedErrorBuilder("MyCustomError").
		Str("reason", "something went wrong").
		Int("code", 500).
		Msg("An error has occurred!")
}
```

### Creating an error using Throw's pre-defined types

Throw provides pre-defined types for common error types, such as `Conflict`, `Forbidden`, `NotFound`, `Unauthorized`,
and `Validation`. These types can be used to create errors with specific HTTP status codes.

```go
package fail

import "github.com/pauloRohling/throw"

func AlwaysFail() error {
	return throw.Internal().
		Str("reason", "something went wrong").
		Int("code", 500).
		Msg("An error has occurred!")
}
```

### Logging a throw error using zerolog

The purpose of Throw is to provide a way to add information to errors in a way that can be helpful for many purposes. In
this example, we will use Throw to create an error and log it using the [rs/zerolog](https://github.com/rs/zerolog)
library. Note that it is possible to extract all the attributes from the error and log them using the `Str()` method.

```go
package fail

import (
	"errors"
	"github.com/pauloRohling/throw"
	"github.com/rs/zerolog/log"
)

func AlwaysFail() error {
	return throw.NewErrorBuilder().
		Str("reason", "something went wrong").
		Int("code", 500).
		Msg("An error has occurred!")
}

func AlwaysFailWithLog() error {
	err := AlwaysFail()
	var throwError *throw.Error
	if errors.As(err, &throwError) {
		LogError(throwError)
	}
	return err
}

func LogError(err *throw.Error) {
	logBuilder := log.Error().Err(err.Unwrap())
	for _, attr := range err.Attributes() {
		logBuilder = logBuilder.Str(attr.Key(), attr.Value())
	}
	logBuilder.Msg(err.Error())
}

```