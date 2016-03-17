package bites

import (
	"encoding/binary"
	"math"
)

func (b Get) GetInt8(i *int8) Get {
	if !b.Space(1) {
		return nil
	}
	*i = int8(b[0])
	return b[1:]
}

func (b Get) GetUint8(i *uint8) Get {
	if !b.Space(1) {
		return nil
	}
	*i = uint8(b[0])
	return b[1:]
}

func (b Get) GetInt16(ip *int16) Get {
	if !b.Space(2) {
		return nil
	}
	*ip = int16(b[0])<<8 + int16(b[1])
	return b[2:]
}
func (b Get) GetInt16LE(ip *int16) Get {
	if !b.Space(2) {
		return nil
	}
	*ip = int16(b[1])<<8 + int16(b[0])
	return b[2:]
}
func (b Get) GetUint16(ip *uint16) Get {
	if !b.Space(2) {
		return nil
	}
	*ip = uint16(b[0])<<8 + uint16(b[1])
	return b[2:]
}
func (b Get) GetUint16LE(ip *uint16) Get {
	if !b.Space(2) {
		return nil
	}
	*ip = uint16(b[1])<<8 + uint16(b[0])
	return b[2:]
}

func (b Get) GetInt32(ip *int32) Get {
	if !b.Space(4) {
		return nil
	}
	*ip = int32(b[0])<<24 + int32(b[1])<<16 + int32(b[2])<<8 + int32(b[3])
	return b[4:]
}
func (b Get) GetInt32LE(ip *int32) Get {
	if !b.Space(4) {
		return nil
	}
	*ip = int32(b[3])<<24 + int32(b[2])<<16 + int32(b[1])<<8 + int32(b[0])
	return b[4:]
}
func (b Get) GetUint32(ip *uint32) Get {
	if !b.Space(4) {
		return nil
	}
	*ip = uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3])
	return b[4:]
}
func (b Get) GetUint32LE(ip *uint32) Get {
	if !b.Space(4) {
		return nil
	}
	*ip = uint32(b[3])<<24 + uint32(b[2])<<16 + uint32(b[1])<<8 + uint32(b[0])
	return b[4:]
}

func (b Get) GetInt64(ip *int64) Get {
	if !b.Space(8) {
		return nil
	}
	*ip = int64(b[0])<<56 + int64(b[1])<<48 + int64(b[2])<<40 + int64(b[3])<<32 + int64(b[4])<<24 + int64(b[5])<<16 + int64(b[6])<<8 + int64(b[7])
	return b[8:]
}
func (b Get) GetInt64LE(ip *int64) Get {
	if !b.Space(8) {
		return nil
	}
	*ip = int64(b[7])<<56 + int64(b[6])<<48 + int64(b[5])<<40 + int64(b[4])<<32 + int64(b[3])<<24 + int64(b[2])<<16 + int64(b[1])<<8 + int64(b[0])
	return b[8:]
}
func (b Get) GetUint64(ip *uint64) Get {
	if !b.Space(8) {
		return nil
	}
	*ip = uint64(b[0])<<56 + uint64(b[1])<<48 + uint64(b[2])<<40 + uint64(b[3])<<32 + uint64(b[4])<<24 + uint64(b[5])<<16 + uint64(b[6])<<8 + uint64(b[7])
	return b[8:]
}
func (b Get) GetUint64LE(ip *uint64) Get {
	if !b.Space(8) {
		return nil
	}
	*ip = uint64(b[7])<<56 + uint64(b[6])<<48 + uint64(b[5])<<40 + uint64(b[4])<<32 + uint64(b[3])<<24 + uint64(b[2])<<16 + uint64(b[1])<<8 + uint64(b[0])
	return b[8:]
}

func (b Get) GetFloat32(f *float32) Get {
	if !b.Space(4) {
		return nil
	}
	var bits uint32
	b = b.GetUint32(&bits)
	*f = math.Float32frombits(bits)
	return b
}

func (b Get) GetFloat64(f *float64) Get {
	if !b.Space(8) {
		return nil
	}
	var bits uint64
	b = b.GetUint64(&bits)
	*f = math.Float64frombits(bits)
	return b
}

func (b Get) GetComplex64(f *complex64) Get {
	if !b.Space(8) {
		return nil
	}
	var r, i float32
	b = b.GetFloat32(&r).GetFloat32(&i)
	*f = complex(r, i)
	return b
}

func (b Get) GetComplex128(f *complex128) Get {
	if !b.Space(16) {
		return nil
	}
	var r, i float64
	b = b.GetFloat64(&r).GetFloat64(&i)
	*f = complex(r, i)
	return b
}

func (b Get) GetVarInt(i *int64, size *int) Get {
	var s int
	*i, s = binary.Varint(b)
	if size != nil {
		*size = s
	}
	return b[s:]
}

func (b Get) GetVarUint(i *uint64, size *int) Get {
	var s int
	*i, s = binary.Uvarint(b)
	if size != nil {
		*size = s
	}
	return b[s:]
}
