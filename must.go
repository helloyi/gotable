package table

func (t *Table) MustInt() int {
	i, err := t.Int()
	if err != nil {
		panic(err)
	}
	return i
}

func (t *Table) MustInt8() int8 {
	i, err := t.Int8()
	if err != nil {
		panic(err)
	}
	return i
}

func (t *Table) MustInt16() int16 {
	i, err := t.Int16()
	if err != nil {
		panic(err)
	}
	return i
}

func (t *Table) MustInt32() int32 {
	i, err := t.Int32()
	if err != nil {
		panic(err)
	}
	return i
}

func (t *Table) MustInt64() int64 {
	i, err := t.Int64()
	if err != nil {
		panic(err)
	}
	return i
}

func (t *Table) MustFloat32() float32 {
	val, err := t.Float32()
	if err != nil {
		panic(err)
	}
	return val
}

func (t *Table) MustFloat64() float64 {
	val, err := t.Float64()
	if err != nil {
		panic(err)
	}
	return val
}

func (t *Table) MustGet(k interface{}) *Table {
	val, err := t.Get(k)
	if err != nil {
		panic(err)
	}
	return val
}
