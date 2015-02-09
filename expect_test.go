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

func TestExpectInts(t *testing.T) {
	b := Empty()
	b = b.PutInt8(-100).PutUint8(200)
	b = b.PutInt16(30000).PutInt16LE(-30000).PutUint16(60000).PutUint16LE(50000)
	b = b.PutInt32(2000000000).PutUint32(4000000000)
	b = b.PutInt32LE(1000000000).PutUint32LE(1234567890)
	b = b.PutInt64(2000000000).PutUint64(4000000000)
	b = b.PutInt64LE(1000000000).PutUint64LE(1234567890)
	b = b.PutFloat32(0.12345).PutFloat64(1.23456)
	b = b.PutComplex64(0.12345 + 1234565i).PutComplex128(1.23456 + 6.2341i)
	b = b.PutVarInt(-91234567789851, nil).PutVarUint(4861749834769813798, nil)

	b = b.ExpectInt8(-100).ExpectUint8(200)
	b = b.ExpectInt16(30000).ExpectInt16LE(-30000).ExpectUint16(60000).ExpectUint16LE(50000)
	b = b.ExpectInt32(2000000000).ExpectUint32(4000000000)
	b = b.ExpectInt32LE(1000000000).ExpectUint32LE(1234567890)
	b = b.ExpectInt64(2000000000).ExpectUint64(4000000000)
	b = b.ExpectInt64LE(1000000000).ExpectUint64LE(1234567890)
	b = b.ExpectFloat32(0.12345).ExpectFloat64(1.23456)
	b = b.ExpectComplex64(0.12345 + 1234565i).ExpectComplex128(1.23456 + 6.2341i)
	b = b.ExpectVarInt(-91234567789851, nil).ExpectVarUint(4861749834769813798, nil)
}
