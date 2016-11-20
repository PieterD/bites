package bites

import "unicode/utf8"

// Bites' purpose is to give byte slices some useful methods.
// The Get methods snip things off the front, and return the remainder of the slice.
// The Expect methods do a Get, and then compare it to the provided value.
// If there is not enough space for a Get, or if the Expect does not match, the method
// will return nil.
// The int methods are big-endian by default, but they have Little-Endian versions too.
// The float and complex methods put them in the form of IEE754 binary representation.
type Get []byte

func (b Get) Put() Put {
	return Put(b)
}

// Return a slice containing the last s bytes.
func (b Get) Last(s int) Get {
	return b[len(b)-s:]
}

// Return a slice without the first s bytes.
func (b Get) Skip(s int) Get {
	return b[s:]
}

// Return a slice with the last s bytes snipped off.
func (b Get) Snip(s int) Get {
	return b[:len(b)-s]
}

// Split the slice into the first s bytes and the rest.
func (b Get) Split(s int) (Get, Get) {
	return b[:s], b[s:]
}

// Make an exact copy of b and return it.
// This will allocate.
func (b Get) Clone() Get {
	clone := make(Get, len(b), len(b))
	copy(clone, b)
	return clone
}

// Return the slice as a string (allocates).
func (b Get) String() string {
	return string(b)
}

// Get one byte, and return the rest.
func (b Get) GetByte(byt *byte) Get {
	*byt = b[0]
	return b[1:]
}

// Get a list of bools.
// The bools are interpreted as though they were written by PutBool.
// If a bool is nil, that bit is skipped.
func (b Get) GetBool(bools ...*bool) Get {
	bytenum := (len(bools)-1)/8 + 1
	if !b.Space(bytenum) {
		return nil
	}
	for i, bol := range bools {
		v := b[i/8] >> uint(7-(i&7)) & 1
		if bol != nil {
			if v == 1 {
				*bol = true
			} else {
				*bol = false
			}
		}
	}
	return b[bytenum:]
}

// Get the first UTF8 rune in b.
// If the rune is not valid UTF8 or b is empty, it returns nil.
func (b Get) GetRune(r *rune, s *int) Get {
	char, size := utf8.DecodeRune(b)
	if char == utf8.RuneError && (size == 0 || size == 1) {
		return nil
	}
	*r = char
	if s != nil {
		*s = size
	}
	return b[size:]
}

// Get a slice of the given size.
func (b Get) GetSlice(slice *[]byte, size int) Get {
	if !b.Space(size) {
		return nil
	}
	*slice = b[:size]
	return b[size:]
}

// Copy b to slice, and return what's left of b.
// If there's not enough in b to fill the slice, return nil.
func (b Get) GetSliceCopy(slice []byte) Get {
	if !b.Space(len(slice)) {
		return nil
	}
	s := copy(slice, b)
	return b[s:]
}

// Get a string of the given size.
// This allocates.
// If there's not enough in b to read the full string, return nil.
func (b Get) GetString(str *string, size int) Get {
	if !b.Space(size) {
		return nil
	}
	var slice []byte
	b = b.GetSlice(&slice, size)
	*str = string(slice)
	return b
}

// Return true if the slice is nil.
// This is poorly named, use More instead.
func (b Get) Error() bool {
	return b == nil
}

// Return true if there is more data to get.
func (b Get) More() bool {
	return b.Len() > 0
}

// Return the length of the slice.
func (b Get) Len() int {
	return len(b)
}

// Returns true if the length of the slice is at least expect.
func (b Get) Space(expect int) bool {
	return b.Len() >= expect
}
