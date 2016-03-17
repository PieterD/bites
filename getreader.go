package bites

import "io"

// A (Byte)Reader that reads everything contained in the given byte slice.
type Reader struct {
	b Get
	n int
}

// A Reader returned by this will allow you to Read everything in the given byte slice.
func (b Get) NewReader() *Reader {
	return &Reader{b: b}
}

// Read a single byte. This is required to implement io.ByteReader.
// When everything has been read, this returns 0, io.EOF.
func (r *Reader) ReadByte() (byte, error) {
	if len(r.b) == 0 {
		return 0, io.EOF
	}
	byt := r.b[0]
	r.b = r.b[1:]
	r.n++
	return byt, nil
}

// Read from the slice. When there is nothing left to read,
// this returns 0, io.EOF.
func (r *Reader) Read(buf []byte) (int, error) {
	l := len(r.b)
	if l == 0 {
		return 0, io.EOF
	}

	if len(buf) < l {
		l = len(buf)
	}
	copy(buf, r.b)
	r.b = (r.b)[l:]
	r.n += l
	return l, nil
}

// Returns the total number of bytes read so far.
func (r *Reader) Total() int {
	return r.n
}
