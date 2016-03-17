package bites

import (
	"fmt"
	"unicode/utf8"
)

type Put []byte

const extendShortLen = 512

var extendShort [extendShortLen]byte

type ErrorInvalidRune rune

func (err ErrorInvalidRune) Error() string {
	return fmt.Sprintf("Invalid rune: %d", err)
}

// Returns an empty slice as a Put.
func New() Put {
	return Put(nil)
}

// Returns b with at least s capacity left.
func (b Put) Capacity(s int) Put {
	if len(b)+s <= cap(b) {
		return b
	}
	orig := len(b)
	b = b.Extend(s)
	return b[:orig]
}

// Return a slice containing the last s bytes.
func (b Put) Last(s int) Put {
	return b[len(b)-s:]
}

// Return a slice with the last s bytes snipped off.
func (b Put) Snip(s int) Put {
	return b[:len(b)-s]
}

// Return the slice as a string (allocates).
func (b Put) String() string {
	return string(b)
}

// Return the slice as a Get.
func (b Put) Get() Get {
	return Get(b)
}

// Extend b by s, return the complete, extended, slice. The extension may contain garbage.
// An allocation (new backing array) will occur if s is larger than cap-len.
func (b Put) Extend(s int) Put {
	l := len(b)
	if l+s <= cap(b) {
		b = b[:l+s]
	} else {
		ob := b
		b = make([]byte, len(b)+s)
		copy(b, ob)
	}
	return b
}

// Set length to 0.
func (b Put) Reuse() Put {
	return b[:0]
}

// Make an exact copy of b and return it.
// This will allocate.
func (b Put) Clone() Put {
	clone := make(Put, len(b), len(b))
	copy(clone, b)
	return clone
}

// Set all bytes to s.
func (b Put) Set(s byte) Put {
	for i := range b {
		b[i] = s
	}
	return b
}

// Set all bytes to 0.
// This is much faster than Set(0).
func (b Put) Zero() Put {
	x := b
	for len(x) > 0 {
		x = x[copy(x, extendShort[:]):]
	}
	return b
}

// Append one byte.
func (b Put) PutByte(byt byte) Put {
	return append(b, byt)
}

// Append a list of bools.
// The bools are set as bits consolidated into bytes,
// and are stored in the same order as they are given.
// This means that if there is a single true given,
// the byte appended will have the value 128.
func (b Put) PutBool(bools ...bool) Put {
	bytenum := (len(bools)-1)/8 + 1
	b = b.Extend(bytenum)
	b.Last(bytenum).Zero()
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

// Append the given rune as UTF8.
// If the rune is not valid, it panics with an error of type ErrorInvalidRune.
func (b Put) PutRune(r rune, s *int) Put {
	l := utf8.RuneLen(r)
	if l == -1 {
		panic(ErrorInvalidRune(r))
	}
	b = b.Extend(l)
	utf8.EncodeRune(b.Last(l), r)
	if s != nil {
		*s = l
	}
	return b
}

// Append the given slice.
func (b Put) PutSlice(slice []byte) Put {
	return append(b, slice...)
}

// Append the given string.
func (b Put) PutString(str string) Put {
	return append(b, str...)
}
