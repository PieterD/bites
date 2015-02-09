package bites

func (b Bites) ExpectInt8(expect int8) Bites {
	var i int8
	b = b.GetInt8(&i)
	if i != expect {
		panic(ErrorExpectInt8{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectUint8(expect uint8) Bites {
	var i uint8
	b = b.GetUint8(&i)
	if i != expect {
		panic(ErrorExpectUint8{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectInt16(expect int16) Bites {
	var i int16
	b = b.GetInt16(&i)
	if i != expect {
		panic(ErrorExpectInt16{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectInt16LE(expect int16) Bites {
	var i int16
	b = b.GetInt16LE(&i)
	if i != expect {
		panic(ErrorExpectInt16{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectUint16(expect uint16) Bites {
	var i uint16
	b = b.GetUint16(&i)
	if i != expect {
		panic(ErrorExpectUint16{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectUint16LE(expect uint16) Bites {
	var i uint16
	b = b.GetUint16LE(&i)
	if i != expect {
		panic(ErrorExpectUint16{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectInt32(expect int32) Bites {
	var i int32
	b = b.GetInt32(&i)
	if i != expect {
		panic(ErrorExpectInt32{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectInt32LE(expect int32) Bites {
	var i int32
	b = b.GetInt32LE(&i)
	if i != expect {
		panic(ErrorExpectInt32{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectUint32(expect uint32) Bites {
	var i uint32
	b = b.GetUint32(&i)
	if i != expect {
		panic(ErrorExpectUint32{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectUint32LE(expect uint32) Bites {
	var i uint32
	b = b.GetUint32LE(&i)
	if i != expect {
		panic(ErrorExpectUint32{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectInt64(expect int64) Bites {
	var i int64
	b = b.GetInt64(&i)
	if i != expect {
		panic(ErrorExpectInt64{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectInt64LE(expect int64) Bites {
	var i int64
	b = b.GetInt64LE(&i)
	if i != expect {
		panic(ErrorExpectInt64{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectUint64(expect uint64) Bites {
	var i uint64
	b = b.GetUint64(&i)
	if i != expect {
		panic(ErrorExpectUint64{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectUint64LE(expect uint64) Bites {
	var i uint64
	b = b.GetUint64LE(&i)
	if i != expect {
		panic(ErrorExpectUint64{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectFloat32(expect float32) Bites {
	var i float32
	b = b.GetFloat32(&i)
	if i != expect {
		panic(ErrorExpectFloat32{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectFloat64(expect float64) Bites {
	var i float64
	b = b.GetFloat64(&i)
	if i != expect {
		panic(ErrorExpectFloat64{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectComplex64(expect complex64) Bites {
	var i complex64
	b = b.GetComplex64(&i)
	if i != expect {
		panic(ErrorExpectComplex64{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectComplex128(expect complex128) Bites {
	var i complex128
	b = b.GetComplex128(&i)
	if i != expect {
		panic(ErrorExpectComplex128{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectVarInt(expect int64, size *int) Bites {
	var i int64
	b = b.GetVarInt(&i, size)
	if i != expect {
		panic(ErrorExpectVarInt{Exp: expect, Got: i})
	}
	return b
}

func (b Bites) ExpectVarUint(expect uint64, size *int) Bites {
	var i uint64
	b = b.GetVarUint(&i, size)
	if i != expect {
		panic(ErrorExpectVarUint{Exp: expect, Got: i})
	}
	return b
}
