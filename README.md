# Throw

[![GoDoc](https://godoc.org/github.com/pauloRohling/throw?status.svg)](https://godoc.org/github.com/pauloRohling/throw)
[![Go Report Card](https://goreportcard.com/badge/github.com/pauloRohling/throw)](https://goreportcard.com/report/github.com/pauloRohling/throw)
[![license](https://img.shields.io/github/license/pauloRohling/throw.svg)](https://github.com/pauloRohling/throw/blob/main/LICENSE)

Throw is a Go library that provides a simple way to build typed errors with attributes.

Throw's API provides a set of chainable methods that allow you to easily build errors with attributes. It also provides
a set of predefined error types that you can use to build your errors.

Throw is inspired by the awesome [rs/zerolog](https://github.com/rs/zerolog) library.

## Installation

```bash
go get github.com/pauloRohling/throw
```

## Features

- Chainable methods for building errors with attributes.
- Predefined error types for common error types.
- Support for standard and advanced attributes.

## Attributes

### Standard Attributes

- `Str`
- `Int`, `Int8`, `Int16`, `Int32`, `Int64`
- `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
- `Float32`, `Float64`

### Advanced Attributes

- `Any`: Formats any type using the `%+v` format.
- `Json`: Formats any type as json using the `json.Marshal` function.

## Examples

Please refer to the [examples](https://github.com/pauloRohling/throw/blob/main/EXAMPLES.md) for more information on how
to use Throw.

## Testing

Throw provides a helper function to assert that the error is of the expected type in a more idiomatic way.

```go
package fail

import "github.com/pauloRohling/throw"

func AlwaysFail() error {
	return throw.NewTypedErrorBuilder("MyCustomError").
		Str("reason", "something went wrong").
		Int("code", 500).
		Msg("An error has occurred!")
}

func TestAlwaysFail(t *testing.T) {
	err := AlwaysFail()
	throw.AssertType(t, err, "MyCustomError")
}
```

## License

Throw is released under the [MIT License](https://github.com/pauloRohling/throw/blob/main/LICENSE).