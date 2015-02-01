package bites

import (
	"bytes"
	"unicode/utf8"
)

type Bites []byte

const extendShortLen = 512

var extendShort [extendShortLen]byte

// Returns an empty slice.
func Empty() Bites {
	return []byte{}
}

// Make a string out of the slice.
// This unavoidably allocates.
func (b Bites) String() string {
	return string(b)
}

// Returns b with at least s capacity left.
// TODO: short path
func (b Bites) Capacity(s int) Bites {
	orig := len(b)
	b = append(b, make([]byte, s)...)
	b = b[:orig]
	return b
}

// Extend b by s, return the complete, extended, slice.
// This may cause an extra allocation if s is much larger than cap-len.
func (b Bites) Extend(s int, zero bool) Bites {
	l := len(b)
	e := l
	if l+s <= cap(b) {
		// Extension fits in cap
		b = b[:l+s]
		e += s
	} else {
		// Extension does not fit, use up all cap first
		b = b[:cap(b)]
		s -= cap(b) - l
		e = len(b)
		if s <= extendShortLen {
			// Short append, alloc-free
			b = append(b, extendShort[:s]...)
			e = len(b)
		} else {
			// Long append, allocates
			b = append(b, make([]byte, s)...)
		}
	}
	if zero {
		x := b[l:e]
		for len(x) > 0 {
			x = x[copy(x, extendShort[:]):]
		}
	}
	return b
}

// Set length to 0.
func (b Bites) Reuse() Bites {
	return b[:0]
}

// Append the given slice.
func (b Bites) PutSlice(slice []byte) Bites {
	return append(b, slice...)
}

// Get a slice of the given size.
func (b Bites) GetSlice(slice *[]byte, size int) Bites {
	*slice = b[:size]
	return b[size:]
}

// Get a slice of the given size.
func (b Bites) GetSliceCopy(slice []byte) Bites {
	return b[copy(slice, b):]
}

// Append the given string.
func (b Bites) PutString(str string) Bites {
	return append(b, str...)
}

// Append one byte.
func (b Bites) PutByte(byt byte) Bites {
	return append(b, byt)
}

// Get one byte.
func (b Bites) GetByte(byt *byte) Bites {
	*byt = b[0]
	return b[1:]
}

// Append the given rune as UTF8.
// If the rune is not valid UTF8, it panics with an ErrorInvalidRune.
func (b Bites) PutRune(r rune) Bites {
	l := utf8.RuneLen(r)
	if l == -1 {
		panic(ErrorInvalidRune(r))
	}
	b = b.Extend(l, false)
	utf8.EncodeRune(b.Last(l), r)
	return b
}

// Make an exact copy and return it.
// This will almost certainly be allocated on the heap.
func (b Bites) Clone() Bites {
	clone := make(Bites, len(b), len(b))
	copy(clone, b)
	return clone
}

// Return a slice containing the last s bytes.
func (b Bites) Last(s int) Bites {
	return b[len(b)-s:]
}

// Return a slice without the first s bytes.
func (b Bites) Skip(s int) Bites {
	return b[s:]
}

// Return a slice with the last s bytes snipped off.
func (b Bites) Snip(s int) Bites {
	return b[:len(b)-s]
}

// Split the slice into the first s bytes and the rest.
func (b Bites) Split(s int) (Bites, Bites) {
	return b[:s], b[s:]
}

// Set all bytes to s.
func (b Bites) Set(s byte) Bites {
	for i := range b {
		b[i] = s
	}
	return b
}

// Set all bytes to 0.
// This is much faster than Set(0).
func (b Bites) Zero() Bites {
	x := b
	for len(x) > 0 {
		x = x[copy(x, extendShort[:]):]
	}
	return b
}

// True if both slices are exactly equal
func (b Bites) Equal(c Bites) bool {
	return bytes.Compare(b, c) == 0
}

// True if the slice is equal to the given string
// TODO: Make sure this doesn't allocate.
func (b Bites) Sequal(str string) bool {
	return b.Equal(Bites(str))
}
