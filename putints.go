package bites

import (
	"encoding/binary"
	"math"
)

func (b Bites) PutInt8(i int8) Bites {
	return append(b, byte(i))
}

func (b Bites) PutUint8(i uint8) Bites {
	return append(b, byte(i))
}

func (b Bites) PutInt16(i int16) Bites {
	return append(b, byte(i>>8), byte(i))
}
func (b Bites) PutInt16LE(i int16) Bites {
	return append(b, byte(i), byte(i>>8))
}
func (b Bites) PutUint16(i uint16) Bites {
	return append(b, byte(i>>8), byte(i))
}
func (b Bites) PutUint16LE(i uint16) Bites {
	return append(b, byte(i), byte(i>>8))
}

func (b Bites) PutInt32(i int32) Bites {
	return append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Bites) PutInt32LE(i int32) Bites {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24))
}
func (b Bites) PutUint32(i uint32) Bites {
	return append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Bites) PutUint32LE(i uint32) Bites {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24))
}

func (b Bites) PutInt64(i int64) Bites {
	return append(b, byte(i>>56), byte(i>>48), byte(i>>40), byte(i>>32), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Bites) PutInt64LE(i int64) Bites {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24), byte(i>>32), byte(i>>40), byte(i>>48), byte(i>>56))
}
func (b Bites) PutUint64(i uint64) Bites {
	return append(b, byte(i>>56), byte(i>>48), byte(i>>40), byte(i>>32), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Bites) PutUint64LE(i uint64) Bites {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24), byte(i>>32), byte(i>>40), byte(i>>48), byte(i>>56))
}

func (b Bites) PutFloat32(f float32) Bites {
	bits := math.Float32bits(f)
	b = b.PutUint32(bits)
	return b
}

func (b Bites) PutFloat64(f float64) Bites {
	bits := math.Float64bits(f)
	b = b.PutUint64(bits)
	return b
}

func (b Bites) PutComplex64(f complex64) Bites {
	r := real(f)
	i := imag(f)
	return b.Capacity(8).PutFloat32(r).PutFloat32(i)
}

func (b Bites) PutComplex128(f complex128) Bites {
	r := real(f)
	i := imag(f)
	return b.Capacity(16).PutFloat64(r).PutFloat64(i)
}

func (b Bites) PutVarInt(i int64, size *int) Bites {
	b = b.Extend(binary.MaxVarintLen64)
	s := binary.PutVarint(b.Last(binary.MaxVarintLen64), i)
	b = b.Snip(binary.MaxVarintLen64 - s)
	if size != nil {
		*size = s
	}
	return b
}

func (b Bites) PutVarUint(i uint64, size *int) Bites {
	b = b.Extend(binary.MaxVarintLen64)
	s := binary.PutUvarint(b.Last(binary.MaxVarintLen64), i)
	b = b.Snip(binary.MaxVarintLen64 - s)
	if size != nil {
		*size = s
	}
	return b
}

func (b Bites) PutVar(ii int) Bites {
	i := int64(ii)
	b = b.Extend(binary.MaxVarintLen64)
	s := binary.PutVarint(b.Last(binary.MaxVarintLen64), i)
	b = b.Snip(binary.MaxVarintLen64 - s)
	return b
}
