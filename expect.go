package bites

import "bytes"

// Read one byte, compare it to byt.
// If it does not match, panic with ErrorExpectByte.
func (b Bites) ExpectByte(byt byte) Bites {
	var bp byte
	b = b.GetByte(&bp)
	if bp != byt {
		panic(ErrorExpectByte{Exp: byt, Got: bp})
	}
	return b
}

//TODO: ExpectBool will require some work to turn bools to bool pointers.

// Read one rune, compare it to byt.
// If it does not match, panic with ErrorExpectRune.
func (b Bites) ExpectRune(r rune) Bites {
	var rp rune
	b = b.GetRune(&rp, nil)
	if rp != r {
		panic(ErrorExpectRune{Exp: r, Got: rp})
	}
	return b
}

// Try to read a slice as long as byt, and see if it matches.
// If it does not, panic with ErrorExpectSlice.
// If there is not enough to read the whole slice, panic with ErrSliceEOF.
func (b Bites) ExpectSlice(byt []byte) Bites {
	var slice []byte
	b = b.GetSlice(&slice, len(byt))
	if bytes.Compare(slice, byt) != 0 {
		panic(ErrorExpectSlice{Exp: byt, Got: slice})
	}
	return b
}

// Try to read a string as long as s, and see if it matches.
// If it does not, panic with ErrorExpectString.
// If it panics with this, it also allocates, but otherwise it does not.
// If there is not enough to read the whole slice, panic with ErrSliceEOF.
func (b Bites) ExpectString(s string) Bites {
	var slice []byte
	b = b.GetSlice(&slice, len(s))
	if !equalByteString(slice, s) {
		panic(ErrorExpectString{Exp: s, Got: string(slice)})
	}
	return b
}

func equalByteString(a []byte, b string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
