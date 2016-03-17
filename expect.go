package bites

import "bytes"

// Read one byte, compare it to byt.
// If it does not match, panic with ErrorExpectByte.
func (b Get) ExpectByte(byt byte) Get {
	var bp byte
	b = b.GetByte(&bp)
	if b.Error() {
		return nil
	}
	if bp != byt {
		return nil
	}
	return b
}

// Try to read a set of bools, and see if it matches.
// See GetBool.
func (b Get) ExpectBool(bools ...bool) Get {
	bo := make([]bool, len(bools))
	bp := make([]*bool, len(bools))
	for i := range bp {
		bp[i] = &bo[i]
	}
	b = b.GetBool(bp...)
	if b.Error() {
		return nil
	}
	for i := range bools {
		if bools[i] != bo[i] {
			return nil
		}
	}
	return b
}

// Read one rune, compare it to byt.
// If it does not match, panic with ErrorExpectRune.
func (b Get) ExpectRune(r rune) Get {
	var rp rune
	b = b.GetRune(&rp, nil)
	if b.Error() {
		return nil
	}
	if rp != r {
		return nil
	}
	return b
}

// Try to read a slice as long as byt, and see if it matches.
// If it does not, panic with ErrorExpectSlice.
// If there is not enough to read the whole slice, panic with ErrSliceEOF.
func (b Get) ExpectSlice(byt []byte) Get {
	var slice []byte
	b = b.GetSlice(&slice, len(byt))
	if b.Error() {
		return nil
	}
	if bytes.Compare(slice, byt) != 0 {
		return nil
	}
	return b
}

// Try to read a string as long as s, and see if it matches.
// If it does not, panic with ErrorExpectString.
// If it panics with this, it also allocates, but otherwise it does not.
// If there is not enough to read the whole slice, panic with ErrSliceEOF.
func (b Get) ExpectString(s string) Get {
	var slice []byte
	b = b.GetSlice(&slice, len(s))
	if b.Error() {
		return nil
	}
	if !equalByteString(slice, s) {
		return nil
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
