package bites

import "testing"

func TestBits8(t *testing.T) {
	b := Empty().PutInt8(100).PutInt8(-100).PutUint8(100).PutUint8(200)
	if b[0] != 100 || b[1] != 256-100 || b[2] != 100 || b[3] != 200 {
		t.Fatalf("Wrong put8 %v", b)
	}

	var sb1, sb2 int8
	var ub1, ub2 uint8

	b.GetInt8(&sb1).GetInt8(&sb2).GetUint8(&ub1).GetUint8(&ub2)

	if sb1 != 100 {
		t.Fatalf("Wrong GetInt8: %d", sb1)
	}
	if sb2 != -100 {
		t.Fatalf("Wrong GetInt8: %d", sb2)
	}
	if ub1 != 100 {
		t.Fatalf("Wrong GetUint8: %d", ub1)
	}
	if ub2 != 200 {
		t.Fatalf("Wrong GetUint8: %d", ub2)
	}
}

func TestBits16(t *testing.T) {
	b := Empty().PutInt16LE(256).PutInt16(1).PutInt16LE(-500).PutUint16LE(256).PutUint16(1)
	if b[0] != 0 || b[1] != 1 {
		t.Fatalf("Wrong put16: %#v", b)
	}
	if b[2] != 0 || b[3] != 1 {
		t.Fatalf("Wrong put16: %#v", b)
	}
	if b[4] != (65536-500)%256 || b[5] != (65536-500)/256 {
		t.Fatalf("Wrong put16: %#v", b)
	}
	if b[6] != 0 || b[7] != 1 {
		t.Fatalf("Wrong put16: %#v", b)
	}
	if b[8] != 0 || b[9] != 1 {
		t.Fatalf("Wrong put16: %#v", b)
	}

	var s1, s2, s3 int16
	var u1, u2 uint16
	b.GetInt16LE(&s1).GetInt16(&s2).GetInt16LE(&s3).GetUint16LE(&u1).GetUint16(&u2)

	if s1 != 256 {
		t.Fatalf("Wrong GetInt16: %d", s1)
	}
	if s2 != 1 {
		t.Fatalf("Wrong GetInt16: %d", s2)
	}
	if s3 != -500 {
		t.Fatalf("Wrong GetInt16: %d", s3)
	}
	if u1 != 256 {
		t.Fatalf("Wrong GetUint16: %d", u1)
	}
	if u2 != 1 {
		t.Fatalf("Wrong GetUint16: %d", u2)
	}
}

func TestBits32(t *testing.T) {
	var s1 int32
	var u1 uint32
	Empty().PutInt32LE(1234567890).PutUint32LE(4123567890).GetInt32LE(&s1).GetUint32LE(&u1)
	if s1 != 1234567890 {
		t.Fatalf("Wrong Put/Get int32: %d", s1)
	}
	if u1 != 4123567890 {
		t.Fatalf("Wrong Put/Get uint32: %d", u1)
	}

	Empty().PutInt32(1234567890).PutUint32(4123567890).GetInt32(&s1).GetUint32(&u1)
	if s1 != 1234567890 {
		t.Fatalf("Wrong Put/Get int32: %d", s1)
	}
	if u1 != 4123567890 {
		t.Fatalf("Wrong Put/Get uint32: %d", u1)
	}
}

func TestBits64(t *testing.T) {
	var s1 int64
	var u1 uint64
	Empty().PutInt64LE(1234567890987654321).PutUint64LE(4123567890).GetInt64LE(&s1).GetUint64LE(&u1)
	if s1 != 1234567890987654321 {
		t.Fatalf("Wrong Put/Get int64: %d", s1)
	}
	if u1 != 4123567890 {
		t.Fatalf("Wrong Put/Get uint64: %d", u1)
	}

	Empty().PutInt64(1234567890).PutUint64(4123567890987654321).GetInt64(&s1).GetUint64(&u1)
	if s1 != 1234567890 {
		t.Fatalf("Wrong Put/Get int64: %d", s1)
	}
	if u1 != 4123567890987654321 {
		t.Fatalf("Wrong Put/Get uint64: %d", u1)
	}
}

func TestFloat(t *testing.T) {
	f1 := float64(0.51234)
	f2 := float32(9124.9876)
	f3 := complex(float32(1.543), float32(9.876))
	f4 := complex(float64(543.12), float64(91235.121))
	var r1 float64
	var r2 float32
	var r3 complex64
	var r4 complex128
	b := Empty().PutFloat64(f1).PutFloat32(f2).PutComplex64(f3).PutComplex128(f4)
	b.GetFloat64(&r1).GetFloat32(&r2).GetComplex64(&r3).GetComplex128(&r4)
	if f1 != r1 || f2 != r2 {
		t.Fatalf("Put/get float failed")
	}
	if f3 != r3 || f4 != f4 {
		t.Fatalf("Put/get complex failed")
	}
}
