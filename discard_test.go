package discard

import (
	"bufio"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func isEOF(err error) bool {
	return err == io.EOF
}

func TestNew(t *testing.T) {
	t.Parallel()
	d := New()
	assert.NotNil(t, d, "New should return a non-nil Discard instance")
}

func TestRead(t *testing.T) {
	t.Parallel()
	d := New()
	n, err := d.Read(make([]byte, 10))
	assert.Equal(t, 0, n, "Read should return 0 bytes read")
	assert.Truef(t, isEOF(err), "Read should return a io.EOF, got: %s", err.Error())
}

func TestReadFrom(t *testing.T) {
	t.Parallel()
	d := New()
	n, err := d.ReadFrom(nil) // Reading from nil should not cause an error
	assert.Equal(t, int64(0), n, "ReadFrom should return 0 bytes read")
	assert.NoError(t, err, "ReadFrom should not return an error")
}

func TestReadAt(t *testing.T) {
	t.Parallel()
	d := New()
	n, err := d.ReadAt(make([]byte, 10), 0)
	assert.Equal(t, 0, n, "ReadAt should return 0 bytes read")
	assert.Truef(t, isEOF(err), "ReadAt should return a io.EOF, got: %s", err.Error())
}

func TestReadByte(t *testing.T) {
	t.Parallel()
	d := New()
	b, err := d.ReadByte()
	assert.Equal(t, byte(0), b, "ReadByte should return 0")
	assert.Truef(t, isEOF(err), "ReadByte should return a io.EOF, got: %s", err.Error())
}

func TestUnreadByte(t *testing.T) {
	t.Parallel()
	d := New()
	err := d.UnreadByte()
	assert.NoError(t, err, "UnreadByte should not return an error")
}

func TestReadRune(t *testing.T) {
	t.Parallel()
	d := New()
	r, size, err := d.ReadRune()
	assert.Equal(t, rune(0), r, "ReadRune should return 0 rune")
	assert.Equal(t, 0, size, "ReadRune should return size 0")
	assert.Truef(t, isEOF(err), "ReadRune should return a io.EOF, got: %s", err.Error())
}

func TestUnreadRune(t *testing.T) {
	t.Parallel()
	d := New()
	err := d.UnreadRune()
	assert.NoError(t, err, "UnreadRune should not return an error")
}

func TestWrite(t *testing.T) {
	t.Parallel()
	d := New()
	n, err := d.Write([]byte("test"))
	assert.Equal(t, 4, n, "Write should return the length of the input")
	assert.NoError(t, err, "Write should not return an error")
}

func TestWriteByte(t *testing.T) {
	t.Parallel()
	d := New()
	err := d.WriteByte('T')
	assert.NoError(t, err, "WriteByte should not return an error")
}

func TestWriteString(t *testing.T) {
	t.Parallel()
	d := New()
	n, err := d.WriteString("test")
	assert.Equal(t, 4, n, "WriteString should return the length of the input string")
	assert.NoError(t, err, "WriteString should not return an error")
}

func TestWriteTo(t *testing.T) {
	t.Parallel()
	d := New()
	n, err := d.WriteTo(nil) // Writing to nil should not cause an error
	assert.Equal(t, int64(0), n, "WriteTo should return 0 bytes written")
	assert.NoError(t, err, "WriteTo should not return an error")
}

func TestSeek(t *testing.T) {
	t.Parallel()
	d := New()
	_, err := d.Seek(0, 0) // Seeking should not cause an error
	assert.NoError(t, err, "Seek should not return an error")
}

func TestWriteAt(t *testing.T) {
	t.Parallel()
	d := New()
	n, err := d.WriteAt([]byte("test"), 0)
	assert.Equal(t, 4, n, "WriteAt should return the length of the input")
	assert.NoError(t, err, "WriteAt should not return an error")
}

func TestClose(t *testing.T) {
	t.Parallel()
	d := New()
	err := d.Close() // Closing should not cause an error
	assert.NoError(t, err, "Close should not return an error")
}

func TestDiscard(t *testing.T) {
	d := New()
	require.NotNil(t, d, "Discard should return a non-nil Discard instance")

	reader := bufio.NewReader(d)
	writer := bufio.NewWriter(d)

	// Test reading from Discard
	_, err := reader.Read(make([]byte, 10))
	assert.Truef(t, isEOF(err), "Reading from Discard should return io.EOF")

	_, err = reader.ReadString('\n')
	assert.Truef(t, isEOF(err), "Reading from Discard with ReadString should return io.EOF")

	// Test writing to Discard
	s := "test"
	n, err := writer.Write([]byte(s))
	assert.Equal(t, len(s), n, "Writing to Discard should return the length of the input")
	assert.NoError(t, err, "Writing to Discard should not return an error")

	err = writer.Flush()
	assert.NoError(t, err, "Flushing Discard writer should not return an error")

	limitR := io.LimitReader(d, 100)
	_, err = limitR.Read(make([]byte, 10))
	assert.Truef(t, isEOF(err), "Reading from LimitReader of Discard should return io.EOF")

	limitW := io.NewOffsetWriter(d, 1)
	n, err = limitW.Write([]byte(s))
	assert.Equal(t, len(s), n, "Writing to OffsetWriter of Discard should return the length of the input")
	assert.NoError(t, err, "Writing to OffsetWriter of Discard should not return an error")

	// Test MultiWriter with Discard
	mult := io.MultiWriter(d, New(), New(), New())
	n, err = mult.Write([]byte(s))
	assert.Equal(t, len(s), n, "Writing to MultiWriter of Discard should return the length of the input")
	assert.NoError(t, err, "Writing to MultiWriter of Discard should not return an error")
}
