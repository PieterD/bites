package bites

// A (Byte)Writer that appends everything written to the given byte slice.
type Writer struct {
	b Bites
	n int
}

// The byte slice provided here is merely as an option to re-use an existing
// slice. When you're done writing to the Writer, call Bites to return an
// updated slice.
func NewWriter(b Bites) *Writer {
	return &Writer{b: b}
}

// Write a single byte. This turns Writer into a ByteWriter.
func (w *Writer) WriteByte(c byte) error {
	w.b = w.b.PutByte(c)
	w.n++
	return nil
}

// This never returns an error.
func (w *Writer) Write(buf []byte) (int, error) {
	w.b = w.b.PutSlice(buf)
	w.n += len(buf)
	return len(buf), nil
}

// Return the total amount of bytes written so far.
func (w *Writer) Total() int {
	return w.n
}

// Return the updated (possibly reallocated) byte slice,
// containing everything written to the Writer.
func (w *Writer) Bites() Bites {
	return w.b
}
