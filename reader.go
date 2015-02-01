package bites

import "io"

type Reader struct {
	b Bites
	n int
}

func (b Bites) Reader() *Reader {
	return &Reader{b: b}
}

func (r *Reader) ReadByte() (byte, error) {
	if len(r.b) == 0 {
		return 0, io.EOF
	}
	byt := r.b[0]
	r.b = r.b[1:]
	r.n++
	return byt, nil
}

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

func (r *Reader) Total() int {
	return r.n
}
