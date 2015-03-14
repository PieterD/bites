package bites

import "fmt"

// Recover from panic.
// If it's an error, put it in *err.
// If it's not an error, continue panicing.
// If no panic was happening, do nothing.
func Catch(err *error) {
	i := recover()
	if i == nil {
		return
	}
	e, ok := i.(error)
	if !ok {
		panic(i)
	}
	*err = e
}

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

type ErrorExpectBool struct {
	Exp, Got []byte
}

func (err ErrorExpectBool) Error() string {
	return fmt.Sprintf("Expected %s (%s), got (%s)", "byte", "bits", "bits")
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

type ErrorExpectInt8 struct {
	Exp, Got int8
}

func (err ErrorExpectInt8) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "int8", err.Exp, err.Got)
}

type ErrorExpectUint8 struct {
	Exp, Got uint8
}

func (err ErrorExpectUint8) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "uint8", err.Exp, err.Got)
}

type ErrorExpectInt16 struct {
	Exp, Got int16
}

func (err ErrorExpectInt16) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "int16", err.Exp, err.Got)
}

type ErrorExpectUint16 struct {
	Exp, Got uint16
}

func (err ErrorExpectUint16) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "uint16", err.Exp, err.Got)
}

type ErrorExpectInt32 struct {
	Exp, Got int32
}

func (err ErrorExpectInt32) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "int32", err.Exp, err.Got)
}

type ErrorExpectUint32 struct {
	Exp, Got uint32
}

func (err ErrorExpectUint32) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "uint32", err.Exp, err.Got)
}

type ErrorExpectInt64 struct {
	Exp, Got int64
}

func (err ErrorExpectInt64) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "int64", err.Exp, err.Got)
}

type ErrorExpectUint64 struct {
	Exp, Got uint64
}

func (err ErrorExpectUint64) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "uint64", err.Exp, err.Got)
}

type ErrorExpectFloat32 struct {
	Exp, Got float32
}

func (err ErrorExpectFloat32) Error() string {
	return fmt.Sprintf("Expected %s (%f), got (%f)", "float32", err.Exp, err.Got)
}

type ErrorExpectFloat64 struct {
	Exp, Got float64
}

func (err ErrorExpectFloat64) Error() string {
	return fmt.Sprintf("Expected %s (%f), got (%f)", "float64", err.Exp, err.Got)
}

type ErrorExpectComplex64 struct {
	Exp, Got complex64
}

func (err ErrorExpectComplex64) Error() string {
	return fmt.Sprintf("Expected %s (%f), got (%f)", "complex64", err.Exp, err.Got)
}

type ErrorExpectComplex128 struct {
	Exp, Got complex128
}

func (err ErrorExpectComplex128) Error() string {
	return fmt.Sprintf("Expected %s (%f), got (%f)", "complex128", err.Exp, err.Got)
}

type ErrorExpectVarInt struct {
	Exp, Got int64
}

func (err ErrorExpectVarInt) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "varint", err.Exp, err.Got)
}

type ErrorExpectVarUint struct {
	Exp, Got uint64
}

func (err ErrorExpectVarUint) Error() string {
	return fmt.Sprintf("Expected %s (%d), got (%d)", "varuint", err.Exp, err.Got)
}
