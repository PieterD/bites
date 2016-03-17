package bites

import (
	"io"
	"testing"
)

func slowCopy(w io.Writer, r io.Reader) (int, error) {
	buf := make([]byte, 2)
	tot := 0
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			return tot, nil
		}
		tot += n
		w.Write(buf[:n])
	}
}

func byteCopy(w io.ByteWriter, r io.ByteReader) (int, error) {
	tot := 0
	for {
		c, err := r.ReadByte()
		if err == io.EOF {
			return tot, nil
		}
		tot++
		w.WriteByte(c)
	}
}

func TestReaderWriterCopy(t *testing.T) {
	b1 := New().PutString("Hello, world!").Get()
	b2 := New().PutString("PFX ")
	r := b1.NewReader()
	w := b2.NewWriter()
	n, err := slowCopy(w, r)
	b2 = w.Bites()
	if n != len("Hello, world!") {
		t.Fatalf("FAIL! Wrong copy length %d", n)
	}
	if n != r.Total() || n != w.Total() {
		t.Fatalf("FAIL! Totals don't match. Expected %d, got reader(%d) and writer(%d)", n, r.Total(), w.Total())
	}
	if err != nil {
		t.Fatalf("FAIL! Unexpected error %v", err)
	}
	if b2.String() != "PFX Hello, world!" {
		t.Fatalf("FAIL! Wrong string after copy append: '%s'", b2.String())
	}

	r = b2.Get().NewReader()
	w = b1.Put().NewWriter()
	n, err = byteCopy(w, r)
	b := w.Bites()
	if n != len("PFX Hello, world!") {
		t.Fatalf("FAIL! Wrong copy length %d", n)
	}
	if n != r.Total() || n != w.Total() {
		t.Fatalf("FAIL! Totals don't match. Expected %d, got reader(%d) and writer(%d)", n, r.Total(), w.Total())
	}
	if err != nil {
		t.Fatalf("FAIL! Unexpected error %v", err)
	}
	if b.String() != "Hello, world!PFX Hello, world!" {
		t.Fatalf("FAIL! Wrong string after copy append: '%s'", b1.String())
	}
}
