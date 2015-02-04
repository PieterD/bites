package bites

import "fmt"

type ErrorInvalidRune rune

func (err ErrorInvalidRune) Error() string {
	return fmt.Sprintf("Invalid rune: %d", err)
}

var ErrSliceEOF error = fmt.Errorf("Unexpectedly reached the end of the bite slice")
var ErrInvalidRune error = fmt.Errorf("Invalid rune")
var ErrInvalidVarintShort error = fmt.Errorf("Invalid varint: buffer too small")
var ErrInvalidVarintOverflow error = fmt.Errorf("Invalid varint: overflow")

type ErrorExpectByte struct {
	Exp, Got byte
}

func (err ErrorExpectByte) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "byte", err.Exp, err.Got)
}

type ErrorExpectRune struct {
	Exp, Got rune
}

func (err ErrorExpectRune) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "rune", err.Exp, err.Got)
}

type ErrorExpectString struct {
	Exp, Got string
}

func (err ErrorExpectString) Error() string {
	return fmt.Sprintf("Expected %s (%s), got (%s)", "string", err.Exp, err.Got)
}

type ErrorExpectSlice struct {
	Exp, Got []byte
}

func (err ErrorExpectSlice) Error() string {
	return fmt.Sprintf("Expected %s (%#v), got (%#v)", "[]byte", err.Exp, err.Got)
}
