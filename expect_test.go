package bites

import "testing"

func TestExpectByteRuneString(t *testing.T) {
	b := New().PutString("Hello 世!")
	b = b.ExpectByte('H').ExpectRune('e').ExpectString("llo ").ExpectRune('世').ExpectByte('!')
	if len(b) != 0 {
		t.Fatalf("FAIL! Expects did not consume the whole string")
	}
}

func TestExpectByteSlice(t *testing.T) {
	slice := []byte("Hello")
	b := New().PutString("Hello world").ExpectSlice(slice)
	if len(b) != 6 {
		t.Fatalf("FAIL! ExpectSlice consumed the wrong amount")
	}
}

func TestExpectFail(t *testing.T) {
	func() {
		defer catch(t, ErrorExpectByte{Exp: 'M', Got: 'm'})
		New().PutString("moo").ExpectByte('M')
	}()
	func() {
		defer catch(t, ErrorExpectRune{Exp: 'M', Got: 'm'})
		New().PutString("moo").ExpectRune('M')
	}()
	func() {
		defer catch(t, ErrorExpectString{Exp: "Str", Got: "str"})
		New().PutString("string").ExpectString("Str")
	}()
	func() {
		defer catch(t, ErrSliceEOF)
		New().PutString("str").ExpectString("String")
	}()
	func() {
		slice := []byte("Hello")
		defer catch(t, ErrorExpectSlice{Exp: slice, Got: []byte("hello")})
		New().PutString("hello world").ExpectSlice(slice)
	}()

	func() {
		defer catch(t, ErrorExpectInt8{Exp: 1, Got: 2})
		New().PutInt8(2).ExpectInt8(1)
	}()
	func() {
		defer catch(t, ErrorExpectUint8{Exp: 1, Got: 2})
		New().PutUint8(2).ExpectUint8(1)
	}()

	func() {
		defer catch(t, ErrorExpectInt16{Exp: 1, Got: 2})
		New().PutInt16(2).ExpectInt16(1)
	}()
	func() {
		defer catch(t, ErrorExpectInt16{Exp: 1, Got: 2})
		New().PutInt16LE(2).ExpectInt16LE(1)
	}()
	func() {
		defer catch(t, ErrorExpectUint16{Exp: 1, Got: 2})
		New().PutUint16(2).ExpectUint16(1)
	}()
	func() {
		defer catch(t, ErrorExpectUint16{Exp: 1, Got: 2})
		New().PutUint16LE(2).ExpectUint16LE(1)
	}()

	func() {
		defer catch(t, ErrorExpectInt32{Exp: 1, Got: 2})
		New().PutInt32(2).ExpectInt32(1)
	}()
	func() {
		defer catch(t, ErrorExpectInt32{Exp: 1, Got: 2})
		New().PutInt32LE(2).ExpectInt32LE(1)
	}()
	func() {
		defer catch(t, ErrorExpectUint32{Exp: 1, Got: 2})
		New().PutUint32(2).ExpectUint32(1)
	}()
	func() {
		defer catch(t, ErrorExpectUint32{Exp: 1, Got: 2})
		New().PutUint32LE(2).ExpectUint32LE(1)
	}()

	func() {
		defer catch(t, ErrorExpectInt64{Exp: 1, Got: 2})
		New().PutInt64(2).ExpectInt64(1)
	}()
	func() {
		defer catch(t, ErrorExpectInt64{Exp: 1, Got: 2})
		New().PutInt64LE(2).ExpectInt64LE(1)
	}()
	func() {
		defer catch(t, ErrorExpectUint64{Exp: 1, Got: 2})
		New().PutUint64(2).ExpectUint64(1)
	}()
	func() {
		defer catch(t, ErrorExpectUint64{Exp: 1, Got: 2})
		New().PutUint64LE(2).ExpectUint64LE(1)
	}()

	func() {
		defer catch(t, ErrorExpectFloat32{Exp: 1, Got: 2})
		New().PutFloat32(2).ExpectFloat32(1)
	}()
	func() {
		defer catch(t, ErrorExpectFloat64{Exp: 1, Got: 2})
		New().PutFloat64(2).ExpectFloat64(1)
	}()

	func() {
		defer catch(t, ErrorExpectComplex64{Exp: 1, Got: 2})
		New().PutComplex64(2).ExpectComplex64(1)
	}()
	func() {
		defer catch(t, ErrorExpectComplex128{Exp: 1, Got: 2})
		New().PutComplex128(2).ExpectComplex128(1)
	}()

	func() {
		defer catch(t, ErrorExpectVarInt{Exp: 1, Got: 2})
		New().PutVarInt(2, nil).ExpectVarInt(1, nil)
	}()
	func() {
		defer catch(t, ErrorExpectVarUint{Exp: 1, Got: 2})
		New().PutVarUint(2, nil).ExpectVarUint(1, nil)
	}()
}

func TestExpectInts(t *testing.T) {
	b := New()
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
