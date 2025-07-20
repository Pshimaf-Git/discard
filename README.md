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

## Examples

### Discarding Output from a Logger

```go
import (
  "log"
  "github.com/Pshimaf-Git/discard"
)

// Create a logger that writes to discard (no output)
logger := log.New(discard.New(), "", log.LstdFlags)
logger.Println("This will not be printed anywhere")
```

### Using as a No-op Writer

```go
import (
  "io"
  "github.com/Pshimaf-Git/discard"
)

func writeSomething(w io.Writer) {
  w.Write([]byte("data"))
}

// Discard output
writeSomething(discard.New())
```

### Using as a No-op Reader

```go
import (
  "io"
  "github.com/Pshimaf-Git/discard"
)

func readSomething(r io.Reader) {
  buf := make([]byte, 10)
  r.Read(buf)
}

// Discard input (always returns EOF)
readSomething(discard.New())
```

### Using as a No-op Closer

```go
import (
  "io"
  "github.com/Pshimaf-Git/discard"
)

func closeResource(c io.Closer) {
  c.Close()
}

// Discard close operation (does nothing)
closeResource(discard.New())
```

You can use `discard.New()` anywhere an `io.Reader`, `io.Writer`, `io.Closer`, or any other `io` interface is required, making it ideal for tests, mocks, or disabling output/input.

## License

MIT License
