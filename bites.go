package bites

import (
	"bytes"
	"unicode/utf8"
)

type Bites []byte

var extendShort [512]byte

// Returns an empty slice.
func Empty() Bites {
	return []byte{}
}

// Make a string out of the slice.
func (b Bites) String() string {
	return string(b)
}

// Returns b with at least s capacity left.
func (b Bites) Capacity(s int) Bites {
	orig := len(b)
	b = append(b, make([]byte, s)...)
	b = b[:orig]
	return b
}

// Extend b by s, return the complete, extended, slice.
// s must be less than or equal to 512, or Extend will panic.
// For larger buffers, use ExtendLong
func (b Bites) Extend(s int) Bites {
	return append(b, extendShort[:s]...)
}

func (b Bites) ExtendLong(s int) Bites {
	return append(b, make([]byte, s)...)
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
	b = b.Extend(l)
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

// Set all bytes to 0.
func (b Bites) Zero() Bites {
	for i := range b {
		b[i] = 0
	}
	return b
}

// True if both slices are exactly equal
func (b Bites) Equal(c Bites) bool {
	return bytes.Compare(b, c) == 0
}

// True if the slice is equal to the given string
func (b Bites) Sequal(str string) bool {
	return b.Equal(Bites(str))
}
