package bites

import (
	"encoding/binary"
	"math"
)

func (b Put) PutInt8(i int8) Put {
	return append(b, byte(i))
}

func (b Put) PutUint8(i uint8) Put {
	return append(b, byte(i))
}

func (b Put) PutInt16(i int16) Put {
	return append(b, byte(i>>8), byte(i))
}
func (b Put) PutInt16LE(i int16) Put {
	return append(b, byte(i), byte(i>>8))
}
func (b Put) PutUint16(i uint16) Put {
	return append(b, byte(i>>8), byte(i))
}
func (b Put) PutUint16LE(i uint16) Put {
	return append(b, byte(i), byte(i>>8))
}

func (b Put) PutInt32(i int32) Put {
	return append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Put) PutInt32LE(i int32) Put {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24))
}
func (b Put) PutUint32(i uint32) Put {
	return append(b, byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Put) PutUint32LE(i uint32) Put {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24))
}

func (b Put) PutInt64(i int64) Put {
	return append(b, byte(i>>56), byte(i>>48), byte(i>>40), byte(i>>32), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Put) PutInt64LE(i int64) Put {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24), byte(i>>32), byte(i>>40), byte(i>>48), byte(i>>56))
}
func (b Put) PutUint64(i uint64) Put {
	return append(b, byte(i>>56), byte(i>>48), byte(i>>40), byte(i>>32), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}
func (b Put) PutUint64LE(i uint64) Put {
	return append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24), byte(i>>32), byte(i>>40), byte(i>>48), byte(i>>56))
}

func (b Put) PutFloat32(f float32) Put {
	bits := math.Float32bits(f)
	b = b.PutUint32(bits)
	return b
}

func (b Put) PutFloat64(f float64) Put {
	bits := math.Float64bits(f)
	b = b.PutUint64(bits)
	return b
}

func (b Put) PutComplex64(f complex64) Put {
	r := real(f)
	i := imag(f)
	return b.Capacity(8).PutFloat32(r).PutFloat32(i)
}

func (b Put) PutComplex128(f complex128) Put {
	r := real(f)
	i := imag(f)
	return b.Capacity(16).PutFloat64(r).PutFloat64(i)
}

func (b Put) PutVarInt(i int64, size *int) Put {
	b = b.Extend(binary.MaxVarintLen64)
	s := binary.PutVarint(b.Last(binary.MaxVarintLen64), i)
	b = b.Snip(binary.MaxVarintLen64 - s)
	if size != nil {
		*size = s
	}
	return b
}

func (b Put) PutVarUint(i uint64, size *int) Put {
	b = b.Extend(binary.MaxVarintLen64)
	s := binary.PutUvarint(b.Last(binary.MaxVarintLen64), i)
	b = b.Snip(binary.MaxVarintLen64 - s)
	if size != nil {
		*size = s
	}
	return b
}

func (b Put) PutVar(ii int) Put {
	i := int64(ii)
	b = b.Extend(binary.MaxVarintLen64)
	s := binary.PutVarint(b.Last(binary.MaxVarintLen64), i)
	b = b.Snip(binary.MaxVarintLen64 - s)
	return b
}
