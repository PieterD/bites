package bites

import (
	"testing"
)

func TestVarint(t *testing.T) {
	var s int
	b := Empty().PutVarInt(1, &s)
	if len(b) != 1 || s != 1 {
		t.Fatalf("Varint size mismatch: got %d, expect 1", s)
	}
	b = b.PutVarInt(-200, &s)
	if len(b) != 3 || s != 2 {
		t.Fatalf("Varint size mismatch: got %d, expect 2", s)
	}
	b = b.PutVarInt(12345, &s)
	if len(b) != 6 || s != 3 {
		t.Fatalf("Varint size mismatch: got %d, expect 3", s)
	}
	b = b.PutVarUint(12345678, &s)
	if len(b) != 10 || s != 4 {
		t.Fatalf("Varint size mismatch: got %d, expect 4", s)
	}
	b = b.PutVarUint(4000000000*4000000000, &s)
	if len(b) != 20 || s != 10 {
		t.Fatalf("Varint size mismatch: got %d, expect 10", s)
	}
	b = b.PutVar(len(b))
	if len(b) != 21 {
		t.Fatalf("Putvar size mismatch")
	}

	var i1, i2, i3 int64
	var i4, i5 uint64
	var i6 int64
	var s1, s2, s3, s4, s5, s6 int
	b.GetVarInt(&i1, &s1).GetVarInt(&i2, &s2).GetVarInt(&i3, &s3).GetVarUint(&i4, &s4).GetVarUint(&i5, &s5).GetVarInt(&i6, &s6)
	if s1 != 1 || s2 != 2 || s3 != 3 || s4 != 4 || s5 != 10 || s6 != 1 {
		t.Fatalf("Invalid sizes %d %d %d %d %d %d", s1, s2, s3, s4, s5, s6)
	}
	if i1 != 1 || i2 != -200 || i3 != 12345 || i4 != 12345678 || i5 != 4000000000*4000000000 || i6 != 20 {
		t.Fatalf("Invalid values")
	}
}
