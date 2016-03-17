package bites

import "testing"

func TestExpectByteRuneString(t *testing.T) {
	b := New().PutString("Hello 世!").Get()
	b = b.ExpectByte('H').ExpectRune('e').ExpectString("llo ").ExpectRune('世').ExpectByte('!')
	if len(b) != 0 {
		t.Fatalf("FAIL! Expects did not consume the whole string")
	}
}

func TestExpectByteSlice(t *testing.T) {
	slice := []byte("Hello")
	b := New().PutString("Hello world").Get().ExpectSlice(slice)
	if len(b) != 6 {
		t.Fatalf("FAIL! ExpectSlice consumed the wrong amount")
	}
}

func TestExpectBool(t *testing.T) {
	b := New().PutBool(true, false, false, true, true, false, false, true, true, false, true).Get()
	if len(b) != 2 {
		t.Fatalf("FAIL! Expected size 2, got %d", len(b))
	}
	b = b.ExpectBool(true, false, false, true, true, false, false, true)
	if len(b) != 1 {
		t.Fatalf("Expected size 1, got %d", len(b))
	}
	b = b.ExpectBool(true, false, true)
	if len(b) != 0 {
		t.Fatalf("Expected size 0, got %d", len(b))
	}
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

	b2 := b.Get()
	b2 = b2.ExpectInt8(-100).ExpectUint8(200)
	b2 = b2.ExpectInt16(30000).ExpectInt16LE(-30000).ExpectUint16(60000).ExpectUint16LE(50000)
	b2 = b2.ExpectInt32(2000000000).ExpectUint32(4000000000)
	b2 = b2.ExpectInt32LE(1000000000).ExpectUint32LE(1234567890)
	b2 = b2.ExpectInt64(2000000000).ExpectUint64(4000000000)
	b2 = b2.ExpectInt64LE(1000000000).ExpectUint64LE(1234567890)
	b2 = b2.ExpectFloat32(0.12345).ExpectFloat64(1.23456)
	b2 = b2.ExpectComplex64(0.12345 + 1234565i).ExpectComplex128(1.23456 + 6.2341i)
	b2 = b2.ExpectVarInt(-91234567789851, nil).ExpectVarUint(4861749834769813798, nil)
}
