# Discard

A lightweight Go package for safely discarding values, errors, or outputs in your code. Useful for ignoring return values or simplifying code where output is intentionally unused.

## Installation

```sh
go get github.com/Pshimaf-Git/discard
```

## Features

- Discard any value or error without side effects.
- Clean up code by removing unnecessary variable assignments.

## Usage

### Discarding Values

```go
import "github.com/Pshimaf-Git/discard"

result := someFunction()
discard.Value(result)
```

### Discarding Errors

```go
import "github.com/Pshimaf-Git/discard"

_, err := anotherFunction()
discard.Error(err)
```

### Discarding Multiple Values

```go
import "github.com/Pshimaf-Git/discard"

a, b, c := getValues()
discard.Values(a, b, c)
```

## License

MIT License
