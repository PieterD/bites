package bites

import "bytes"

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
func (b Bites) Capacity(s int) Bites {
	if len(b)+s <= cap(b) {
		return b
	}
	orig := len(b)
	b = b.Extend(s, false)
	return b[:orig]
}

// Extend b by s, return the complete, extended, slice.
// If zero is true, the extension bytes are set to 0, otherwise their content is left as-is.
// An allocation (new backing array) will occur if s is larger than cap-len.
// An extra allocation (for a temporary slice) will occur if s is much larger (by 512 bytes) than cap-len.
func (b Bites) Extend(s int, zero bool) Bites {
	l := len(b)
	e := l
	if l+s <= cap(b) {
		// Short append; extension fits in cap, no allocation.
		b = b[:l+s]
		e += s
	} else {
		// Extension does not fit, use up all cap first
		e = cap(b)
		b = b[:e]
		s -= e - l
		if s <= extendShortLen {
			// Mid append, must allocate new backing array.
			b = append(b, extendShort[:s]...)
		} else {
			// Long append, must allocate new backing array and temporary slice.
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

// Make an exact copy of b and return it.
// This will allocate.
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
