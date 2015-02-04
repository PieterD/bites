package bites

import "testing"

func TestExpectByteRuneString(t *testing.T) {
	b := Empty().PutString("Hello 世!")
	b = b.ExpectByte('H').ExpectRune('e').ExpectString("llo ").ExpectRune('世').ExpectByte('!')
	if len(b) != 0 {
		t.Fatalf("FAIL! Expects did not consume the whole string")
	}
}

func TestExpectByteSlice(t *testing.T) {
	slice := []byte("Hello")
	b := Empty().PutString("Hello world").ExpectSlice(slice)
	if len(b) != 6 {
		t.Fatalf("FAIL! ExpectSlice consumed the wrong amount")
	}
}

func TestExpectFail(t *testing.T) {
	func() {
		defer catch(t, ErrorExpectByte{Exp: 'M', Got: 'm'})
		Empty().PutString("moo").ExpectByte('M')
	}()
	func() {
		defer catch(t, ErrorExpectRune{Exp: 'M', Got: 'm'})
		Empty().PutString("moo").ExpectRune('M')
	}()
	func() {
		defer catch(t, ErrorExpectString{Exp: "Str", Got: "str"})
		Empty().PutString("string").ExpectString("Str")
	}()
	func() {
		defer catch(t, ErrSliceEOF)
		Empty().PutString("str").ExpectString("String")
	}()
	func() {
		slice := []byte("Hello")
		defer catch(t, ErrorExpectSlice{Exp: slice, Got: []byte("hello")})
		Empty().PutString("hello world").ExpectSlice(slice)
	}()
}
