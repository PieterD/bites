package bites

import "encoding/binary"

func (b Bites) GetInt8(i *int8) Bites {
	*i = int8(b[0])
	return b[1:]
}

func (b Bites) GetUint8(i *uint8) Bites {
	*i = uint8(b[0])
	return b[1:]
}

func (b Bites) GetInt16(ip *int16) Bites {
	*ip = int16(b[0])<<8 + int16(b[1])
	return b[2:]
}
func (b Bites) GetInt16LE(ip *int16) Bites {
	*ip = int16(b[1])<<8 + int16(b[0])
	return b[2:]
}
func (b Bites) GetUint16(ip *uint16) Bites {
	*ip = uint16(b[0])<<8 + uint16(b[1])
	return b[2:]
}
func (b Bites) GetUint16LE(ip *uint16) Bites {
	*ip = uint16(b[1])<<8 + uint16(b[0])
	return b[2:]
}

func (b Bites) GetInt32(ip *int32) Bites {
	*ip = int32(b[0])<<24 + int32(b[1])<<16 + int32(b[2])<<8 + int32(b[3])
	return b[4:]
}
func (b Bites) GetInt32LE(ip *int32) Bites {
	*ip = int32(b[3])<<24 + int32(b[2])<<16 + int32(b[1])<<8 + int32(b[0])
	return b[4:]
}
func (b Bites) GetUint32(ip *uint32) Bites {
	*ip = uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3])
	return b[4:]
}
func (b Bites) GetUint32LE(ip *uint32) Bites {
	*ip = uint32(b[3])<<24 + uint32(b[2])<<16 + uint32(b[1])<<8 + uint32(b[0])
	return b[4:]
}

func (b Bites) GetInt64(ip *int64) Bites {
	*ip = int64(b[0])<<56 + int64(b[1])<<48 + int64(b[2])<<40 + int64(b[3])<<32 + int64(b[4])<<24 + int64(b[5])<<16 + int64(b[6])<<8 + int64(b[7])
	return b[8:]
}
func (b Bites) GetInt64LE(ip *int64) Bites {
	*ip = int64(b[7])<<56 + int64(b[6])<<48 + int64(b[5])<<40 + int64(b[4])<<32 + int64(b[3])<<24 + int64(b[2])<<16 + int64(b[1])<<8 + int64(b[0])
	return b[8:]
}
func (b Bites) GetUint64(ip *uint64) Bites {
	*ip = uint64(b[0])<<56 + uint64(b[1])<<48 + uint64(b[2])<<40 + uint64(b[3])<<32 + uint64(b[4])<<24 + uint64(b[5])<<16 + uint64(b[6])<<8 + uint64(b[7])
	return b[8:]
}
func (b Bites) GetUint64LE(ip *uint64) Bites {
	*ip = uint64(b[7])<<56 + uint64(b[6])<<48 + uint64(b[5])<<40 + uint64(b[4])<<32 + uint64(b[3])<<24 + uint64(b[2])<<16 + uint64(b[1])<<8 + uint64(b[0])
	return b[8:]
}

func (b Bites) GetVarInt(i *int64, size *int) Bites {
	var s int
	*i, s = binary.Varint(b)
	if size != nil {
		*size = s
	}
	return b[s:]
}

func (b Bites) GetVarUint(i *uint64, size *int) Bites {
	var s int
	*i, s = binary.Uvarint(b)
	if size != nil {
		*size = s
	}
	return b[s:]
}
