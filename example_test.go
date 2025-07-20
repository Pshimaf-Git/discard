package discard

func ExampleDiscard() {
	// Create a new Discard instance
	d := New()

	// Use the Discard instance to read and write data
	n, err := d.Read(make([]byte, 10))
	if n != 0 || err != nil {
		panic("Read should return 0 bytes and no error")
	}

	n, err = d.Write([]byte("Hello"))
	if n != 5 || err != nil {
		panic("Write should return 5 bytes and no error")
	}

	// ReadFrom should also return 0 bytes and no error
	n64, err := d.ReadFrom(nil)
	if n64 != 0 || err != nil {
		panic("ReadFrom should return 0 bytes and no error")
	}

	// ReadAt should return 0 bytes and no error
	n, err = d.ReadAt(make([]byte, 10), 0)
	if n != 0 || err != nil {
		panic("ReadAt should return 0 bytes and no error")
	}

	// ReadByte should return 0 and no error
	b, err := d.ReadByte()
	if b != 0 || err != nil {
		panic("ReadByte should return 0 and no error")
	}

	// UnreadByte should not return an error
	err = d.UnreadByte()
	if err != nil {
		panic("UnreadByte should not return an error")
	}

	// Output:
}
