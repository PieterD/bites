package bites

import "unicode/utf8"

// Append one byte.
func (b Bites) PutByte(byt byte) Bites {
	return append(b, byt)
}

// Get one byte, and return the rest.
func (b Bites) GetByte(byt *byte) Bites {
	*byt = b[0]
	return b[1:]
}

// Append a list of bools.
// The bools are set as bits consolidated into bytes,
// and are stored in the same order as they are given.
// This means that if there is a single true given,
// the byte appended will have the value 128.
func (b Bites) PutBool(bools ...bool) Bites {
	bytenum := len(bools)/8 + 1
	b = b.Extend(bytenum, true)
	bts := b.Last(bytenum)
	for i, bol := range bools {
		v := byte(0)
		if bol {
			v = 1
		}
		v = v << uint(7-(i&7))
		bts[i/8] = bts[i/8] | v
	}
	return b
}

// Get a list of bools.
// The bools are interpreted as though they were written by PutBool.
// If a bool is nil, that bit is skipped.
func (b Bites) GetBool(bools ...*bool) Bites {
	bytenum := len(bools)/8 + 1
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

// Append the given rune as UTF8.
// If the rune is not valid UTF8, it panics with an error of type ErrorInvalidRune.
func (b Bites) PutRune(r rune, s *int) Bites {
	l := utf8.RuneLen(r)
	if l == -1 {
		panic(ErrorInvalidRune(r))
	}
	b = b.Extend(l, false)
	utf8.EncodeRune(b.Last(l), r)
	if s != nil {
		*s = l
	}
	return b
}

// Get the first UTF8 rune in b.
// If the rune is not valid UTF8 or b is empty, it panics with ErrInvalidRune.
func (b Bites) GetRune(r *rune, s *int) Bites {
	char, size := utf8.DecodeRune(b)
	if char == utf8.RuneError && (size == 0 || size == 1) {
		panic(ErrInvalidRune)
	}
	*r = char
	if s != nil {
		*s = size
	}
	return b[size:]
}

// Append the given slice.
func (b Bites) PutSlice(slice []byte) Bites {
	return append(b, slice...)
}

// Append the given string.
func (b Bites) PutString(str string) Bites {
	return append(b, str...)
}

// Get a slice of the given size.
func (b Bites) GetSlice(slice *[]byte, size int) Bites {
	*slice = b[:size]
	return b[size:]
}

// Copy b to slice, and return what's left of b.
func (b Bites) GetSliceCopy(slice []byte) Bites {
	return b[copy(slice, b):]
}
