package bites

import "fmt"

type ErrorInvalidRune rune

func (err ErrorInvalidRune) Error() string {
	return fmt.Sprintf("Invalid rune: %d", err)
}

var ErrInvalidRune error = fmt.Errorf("Invalid rune")
var ErrInvalidVarintShort error = fmt.Errorf("Invalid varint: buffer too small")
var ErrInvalidVarintOverflow error = fmt.Errorf("Invalid varint: overflow")
