package bites

func (b Get) ExpectInt8(expect int8) Get {
	var i int8
	b = b.GetInt8(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectUint8(expect uint8) Get {
	var i uint8
	b = b.GetUint8(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectInt16(expect int16) Get {
	var i int16
	b = b.GetInt16(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectInt16LE(expect int16) Get {
	var i int16
	b = b.GetInt16LE(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectUint16(expect uint16) Get {
	var i uint16
	b = b.GetUint16(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectUint16LE(expect uint16) Get {
	var i uint16
	b = b.GetUint16LE(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectInt32(expect int32) Get {
	var i int32
	b = b.GetInt32(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectInt32LE(expect int32) Get {
	var i int32
	b = b.GetInt32LE(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectUint32(expect uint32) Get {
	var i uint32
	b = b.GetUint32(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectUint32LE(expect uint32) Get {
	var i uint32
	b = b.GetUint32LE(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectInt64(expect int64) Get {
	var i int64
	b = b.GetInt64(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectInt64LE(expect int64) Get {
	var i int64
	b = b.GetInt64LE(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectUint64(expect uint64) Get {
	var i uint64
	b = b.GetUint64(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectUint64LE(expect uint64) Get {
	var i uint64
	b = b.GetUint64LE(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectFloat32(expect float32) Get {
	var i float32
	b = b.GetFloat32(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectFloat64(expect float64) Get {
	var i float64
	b = b.GetFloat64(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectComplex64(expect complex64) Get {
	var i complex64
	b = b.GetComplex64(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectComplex128(expect complex128) Get {
	var i complex128
	b = b.GetComplex128(&i)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectVarInt(expect int64, size *int) Get {
	var i int64
	b = b.GetVarInt(&i, size)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}

func (b Get) ExpectVarUint(expect uint64, size *int) Get {
	var i uint64
	b = b.GetVarUint(&i, size)
	if b.Error() {
		return nil
	}
	if i != expect {
		return nil
	}
	return b
}
