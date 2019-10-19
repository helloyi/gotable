package table

// MustInt must api for Int()
func (t *Table) MustInt() int {
	i, err := t.Int()
	if err != nil {
		panic(err)
	}
	return i
}

// MustInt8 must api for Int8()
func (t *Table) MustInt8() int8 {
	i, err := t.Int8()
	if err != nil {
		panic(err)
	}
	return i
}

// MustInt16 must api for Int16()
func (t *Table) MustInt16() int16 {
	i, err := t.Int16()
	if err != nil {
		panic(err)
	}
	return i
}

// MustInt32 must api for Int32()
func (t *Table) MustInt32() int32 {
	i, err := t.Int32()
	if err != nil {
		panic(err)
	}
	return i
}

// MustInt64 must api for Int64()
func (t *Table) MustInt64() int64 {
	i, err := t.Int64()
	if err != nil {
		panic(err)
	}
	return i
}

// MustFloat32 must api for Float32()
func (t *Table) MustFloat32() float32 {
	val, err := t.Float32()
	if err != nil {
		panic(err)
	}
	return val
}

// MustFloat64 must api for Float64()
func (t *Table) MustFloat64() float64 {
	val, err := t.Float64()
	if err != nil {
		panic(err)
	}
	return val
}

// MustGet must api for Get
func (t *Table) MustGet(k interface{}) *Table {
	val, err := t.Get(k)
	if err != nil {
		panic(err)
	}
	return val
}

// MustMap must api for Map
func (t *Table) MustMap() map[*Table]*Table {
	tm, err := t.Map()
	if err != nil {
		panic(err)
	}
	return tm
}

// MustSlice must api for Slice
func (t *Table) MustSlice() []*Table {
	ts, err := t.Slice()
	if err != nil {
		panic(err)
	}
	return ts
}

// MustAList must api for AList
func (t *Table) MustAList() [][2]*Table {
	tl, err := t.AList()
	if err != nil {
		panic(err)
	}
	return tl
}

// MustPList must api for PList
func (t *Table) MustPList() []*Table {
	tl, err := t.PList()
	if err != nil {
		panic(err)
	}
	return tl
}
