package bites

import "testing"

func BenchmarkExtendShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bts := make(Bites, 50)
		bts = bts.Reuse().Extend(40)
	}
}

func BenchmarkExtendShortReuse(b *testing.B) {
	bts := make(Bites, 50)
	for i := 0; i < b.N; i++ {
		bts = bts.Reuse().Extend(40)
	}
}

func BenchmarkExtendMid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bts := make(Bites, 50)
		bts = bts.Reuse().Extend(400)
	}
}

func BenchmarkExtendMidReuse(b *testing.B) {
	bts := make(Bites, 50)
	for i := 0; i < b.N; i++ {
		bts = bts.Reuse().Extend(400)
	}
}

func BenchmarkExtendLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bts := make(Bites, 50)
		bts = bts.Reuse().Extend(4000)
	}
}

func BenchmarkExtendLongReuse(b *testing.B) {
	bts := make(Bites, 50)
	for i := 0; i < b.N; i++ {
		bts = bts.Reuse().Extend(4000)
	}
}

func BenchmarkEverything(b *testing.B) {
	bts := Empty()
	for i := 0; i < b.N; i++ {
		bts = doStuff(bts, b)
	}
}

type failer interface {
	Fail()
}

func TestEverything(t *testing.T) {
	bts := Empty()
	doStuff(bts, t)
}

func doStuff(bts Bites, f failer) Bites {
	var hash [32]byte
	//return bts.Reuse().PutString("hello").PutSlice(hash[:]).PutByte(4).PutByte(4).PutVar(5)
	bts = bts.Reuse()
	bts = bts.PutString("hello")
	bts = bts.PutSlice(hash[:])
	bts = bts.PutByte(4)
	bts = bts.PutByte(4)
	bts = bts.Extend(10)
	bts = bts.PutInt16(12)
	bts = bts.PutUint16(250)
	bts = bts.PutInt16LE(-12)
	bts = bts.PutUint16LE(2)
	bts = bts.PutInt32(2000000000)
	bts = bts.PutUint32LE(4000000000)
	bts = bts.PutUint64LE(3000000000 * 4000000000)
	bts = bts.PutInt64(2000000000 * 1000000000)
	bts = bts.PutVar(1)
	bts = bts.PutVar(1000000000)
	bts = bts.PutComplex64(987.654 + 123.456i)
	bts = bts.PutComplex128(123.456 + 987.654i)

	rv := bts

	var helloSlice []byte
	var b1, b2 byte
	var s16 int16
	var u16 uint16
	var s32 int32
	var u32 uint32
	var u64 uint64
	var s64, v64 int64
	var c64 complex64
	var c128 complex128
	var size int
	bts = bts.GetSlice(&helloSlice, 5)
	bts = bts.GetSliceCopy(hash[:])
	bts = bts.GetByte(&b1)
	if b1 != 4 {
		f.Fail()
	}
	bts = bts.GetByte(&b2)
	if b2 != 4 {
		f.Fail()
	}
	bts = bts.Skip(10)
	bts = bts.GetInt16(&s16)
	if s16 != 12 {
		f.Fail()
	}
	bts = bts.GetUint16(&u16)
	if u16 != 250 {
		f.Fail()
	}
	bts = bts.GetInt16LE(&s16)
	if s16 != -12 {
		f.Fail()
	}
	bts = bts.GetUint16LE(&u16)
	if u16 != 2 {
		f.Fail()
	}
	bts = bts.GetInt32(&s32)
	if s32 != 2000000000 {
		f.Fail()
	}
	bts = bts.GetUint32LE(&u32)
	if u32 != 4000000000 {
		f.Fail()
	}
	bts = bts.GetUint64LE(&u64)
	if u64 != 3000000000*4000000000 {
		f.Fail()
	}
	bts = bts.GetInt64(&s64)
	if s64 != 2000000000*1000000000 {
		f.Fail()
	}
	bts = bts.GetVarInt(&v64, &size)
	if v64 != 1 || size != 1 {
		f.Fail()
	}
	bts = bts.GetVarInt(&v64, &size)
	if v64 != 1000000000 || size != 5 {
		f.Fail()
	}
	bts = bts.GetComplex64(&c64)
	if c64 != 987.654+123.456i {
		f.Fail()
	}
	bts = bts.GetComplex128(&c128)
	if c128 != 123.456+987.654i {
		f.Fail()
	}
	return rv
}
