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

## Attributes

### Standard Attributes

- `Str`
- `Int`, `Int8`, `Int16`, `Int32`, `Int64`
- `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
- `Float32`, `Float64`

### Advanced Attributes

- `Any`: Formats any type using the `%+v` format.
- `Json`: Formats any type as json using the `json.Marshal` function.

## License

Throw is released under the [MIT License](https://github.com/pauloRohling/throw/blob/main/LICENSE).