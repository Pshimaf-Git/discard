# Discard

## Overview

This package provides a type that implements all interfaces from Go's `io` package. It is designed to serve as a stub or placeholder wherever any `io` interface is required. This can be useful for testing, scaffolding, or situations where a no-op implementation is needed.

## Installation

```sh
go get github.com/Pshimaf-Git/discard
```

## Features

- Implements all interfaces from the `io` package (e.g., `io.Reader`, `io.Writer`, `io.Closer`, etc.).
- Can be used as a drop-in replacement in any context expecting an `io` interface.
- Useful for testing and mocking I/O operations.

## Usage

```go
import
(
   "log"

   "github.com/Pshimaf-Git/discard"
)

// for tests use logger which does not output anything
var discardLogger = log.New(discard.New(), "", 1)

var myhandler = NewMyHandler(db, cache, discardLogger)

// testing code ...
```

## License

MIT License
