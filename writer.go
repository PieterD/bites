package bites

type Writer struct {
	b Bites
	n int
}

func (b Bites) Writer() *Writer {
	return &Writer{b: b}
}

func (w *Writer) WriteByte(c byte) error {
	w.b = w.b.PutByte(c)
	w.n++
	return nil
}

func (w *Writer) Write(buf []byte) (int, error) {
	w.b = w.b.PutSlice(buf)
	w.n += len(buf)
	return len(buf), nil
}

func (w *Writer) Total() int {
	return w.n
}

func (w *Writer) Bites() Bites {
	return w.b
}
