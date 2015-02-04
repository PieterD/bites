package bites

import (
	"reflect"
	"testing"
)

func TestExpectByteRuneString(t *testing.T) {
	b := Empty().PutString("Hello 世!")
	b = b.ExpectByte('H').ExpectRune('e').ExpectString("llo ").ExpectRune('世').ExpectByte('!')
	if len(b) != 0 {
		t.Fatalf("FAIL! Expects did not consume the whole string")
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
}

func catch(t *testing.T, expect error) {
	e := recover()
	if e == nil {
		t.Fatalf("Expected panic: %#v", expect)
		return
	}
	got, ok := e.(error)
	if !ok {
		panic(e)
	}

	if reflect.TypeOf(got) != reflect.TypeOf(expect) {
		t.Fatalf("Expected panic %#v, got %#v", expect, got)
	}
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("Expected panic %#v, got %#v", expect, got)
	}
}
