// Package discard provides a Discard type that implements various io interfaces.
// It is used to discard any data read or written, effectively acting as a no-op.
package discard

import (
	"io"
)

// Discard is a type that implements the io.Reader, io.Writer, and other interfaces.
// It is used to discard any data read or written, effectively acting as a no-op.
type Discard struct{}

// Ensure that Discard implements the necessary interfaces.
// This is a compile-time assertion to ensure that Discard satisfies the interfaces.
var (
	_ io.Reader          = Discard{}
	_ io.ReaderAt        = Discard{}
	_ io.ReaderFrom      = Discard{}
	_ io.ByteReader      = Discard{}
	_ io.RuneReader      = Discard{}
	_ io.Writer          = Discard{}
	_ io.WriterAt        = Discard{}
	_ io.WriterTo        = Discard{}
	_ io.ByteWriter      = Discard{}
	_ io.StringWriter    = Discard{}
	_ io.Seeker          = Discard{}
	_ io.Closer          = Discard{}
	_ io.ReadCloser      = Discard{}
	_ io.WriteCloser     = Discard{}
	_ io.ReadSeeker      = Discard{}
	_ io.ReadWriter      = Discard{}
	_ io.WriteSeeker     = Discard{}
	_ io.StringWriter    = Discard{}
	_ io.ByteScanner     = Discard{}
	_ io.RuneScanner     = Discard{}
	_ io.ReaderFrom      = Discard{}
	_ io.WriterAt        = Discard{}
	_ io.ReaderAt        = Discard{}
	_ io.WriterTo        = Discard{}
	_ io.ReadWriteCloser = Discard{}
	_ io.ReadWriteSeeker = Discard{}
)

// New creates and returns a new Discard instance that implements the all interfaces from io library.
// This instance can be used wherever an [io.Reader], [io.Writer], or other io interfaces are required,
// effectively discarding any data read or written.
// It is useful for testing or when you need a no-op implementation of these interfaces.
func New() Discard {
	return Discard{}
}

// Read implements the [io.Reader] interface for Discard.
// It discards any input and always returns 0, [io.EOF], indicating that no bytes were read and no error occurred.
func (Discard) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

// ReadFrom implements the [io.ReaderFrom] interface for Discard.
// It reads from the provided [io.Reader] and discards all data, always returning 0, nil.
func (Discard) ReadFrom(r io.Reader) (n int64, err error) {
	return
}

// ReadAt implements the [io.ReaderAt] interface for Discard.
// It discards any input and always returns 0, [io.EOF].
func (Discard) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, io.EOF
}

// ReadByte implements the [io.ByteReader] interface for Discard.
// It always returns 0, [io.EOF], discarding any input.
func (Discard) ReadByte() (b byte, err error) {
	return 0, io.EOF
}

// UnreadByte implements the [io.ByteScanner] interface for Discard.
// It always returns nil, discarding any input.
func (Discard) UnreadByte() (err error) {
	return
}

// ReadRune implements the [io.RuneReader] interface for Discard.
// It always returns 0, 0, [io.EOF], discarding any input.
func (Discard) ReadRune() (r rune, size int, err error) {
	return 0, 0, io.EOF
}

// UnreadRune implements the [io.RuneScanner] interface for Discard.
// It always returns nil, discarding any input.
func (Discard) UnreadRune() (err error) {
	return
}

// Write implements the [io.Writer] interface for Discard.
// It discards any data written to it and always returns len(p), nil.
func (Discard) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// WriteByte implements the [io.ByteWriter] interface for Discard.
// It discards the byte and always returns nil.
func (Discard) WriteByte(b byte) (err error) {
	return
}

// WriteString implements the [io.StringWriter] interface for Discard.
// It discards the string and always returns len(s), nil.
func (Discard) WriteString(s string) (n int, err error) {
	return len(s), nil
}

// WriteTo implements the [io.WriterTo] interface for Discard.
// It writes nothing to the provided [io.Writer] and always returns 0, nil.
func (Discard) WriteTo(w io.Writer) (n int64, err error) {
	return
}

// WriteAt implements the [io.WriterAt] interface for Discard.
// It discards any data written to it and always returns len(p), nil.
func (Discard) WriteAt(p []byte, off int64) (n int, err error) {
	return len(p), nil
}

// Seek implements the [io.Seeker] interface for Discard.
// It always returns 0, nil, as there is no underlying data to seek.
func (Discard) Seek(offset int64, whence int) (n int64, err error) {
	return
}

// Close implements the [io.Closer] interface for Discard.
// It always returns nil, as there are no resources to release.
func (Discard) Close() (err error) {
	return
}
