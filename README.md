# README.md

## Overview

This package provides a type that implements all interfaces from Go's `io` package. It is designed to serve as a stub or placeholder wherever any `io` interface is required. This can be useful for testing, scaffolding, or situations where a no-op implementation is needed.

## InstallInstallation

```sh
go get github.com/Pshimaf-Git/discard
```

## Features

- Implements all interfaces from the `io` package (e.g., `io.Reader`, `io.Writer`, `io.Closer`, etc.).
- Can be used as a drop-in replacement in any context expecting an `io` interface.
- Useful for testing and mocking I/O operations.

## Usage

```go
import "github.com/Pshimaf-Git/discard"

var discard io.Reader = discard.New()
```

## License

MIT License
