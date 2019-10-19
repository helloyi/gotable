package table

import (
	"reflect"
)

func (t *Table) getv() reflect.Value {
	if t.v.Kind() == reflect.Invalid {
		t.v = reflect.ValueOf(t.i)
	}

	return t.v
}

func (t *Table) geti() interface{} {
	if t.i == nil {
		t.i = t.v.Interface()
	}
	return t.i
}

//// get op

func (t *Table) mapGet(k interface{}) *Table {
	v := t.getv().MapIndex(reflect.ValueOf(k))
	if v.Kind() == reflect.Invalid {
		return nil
	}
	return &Table{v: v}
}

func (t *Table) sliceGet(idx int) *Table {
	l := t.getv().Len()
	if idx >= l {
		return nil
	}

	v := t.getv().Index(idx)
	return &Table{v: v}
}

func (t *Table) structGet(field string) *Table {
	v := t.getv().FieldByName(field)
	if v.Kind() == reflect.Invalid {
		return nil
	}
	return &Table{v: v}
}

//// put op

func (t *Table) mapPut(k, v interface{}) error {
	if t.getv().IsNil() {
		t.v = reflect.MakeMap(t.getv().Type())
		t.i = t.v.Interface()
	}
	t.getv().SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	return nil
}

func (t *Table) arrayPut(idx int, v interface{}) error {
	cap := t.getv().Cap()
	if idx >= cap {
		return nil
	}
	ev := t.getv().Index(idx)
	ev.Set(reflect.ValueOf(v))
	return nil
}

func (t *Table) slicePut(idx int, v interface{}) error {
	l := t.getv().Len()
	if idx < l { // set
		ev := t.getv().Index(idx)
		ev.Set(reflect.ValueOf(v))
	} else { // append
		sv := t.getv()
		x := reflect.ValueOf(v)
		zv := reflect.Zero(x.Type())
		for i := l; i < idx; i++ {
			sv = reflect.Append(sv, zv)
		}
		sv = reflect.Append(sv, x)

		t.v = sv
		t.i = sv.Interface()
	}
	return nil
}

// structPut ...
func (t *Table) structPut(fn string, v interface{}) error {
	fv := t.getv().FieldByName(fn)
	if fv.IsValid() {
		return &ErrNotExist{"Table.structPut", fn + " field"}
	}
	fv.Set(reflect.ValueOf(v))
	return nil
}

func (t *Table) bool() bool {
	return t.getv().Bool()
}

func (t *Table) int() int64 {
	return t.getv().Int()
}

func (t *Table) uint() uint64 {
	return t.getv().Uint()
}

func (t *Table) float() float64 {
	return t.getv().Float()
}

func (t *Table) complex_() complex128 {
	return t.getv().Complex()
}

func (t *Table) string() string {
	return t.getv().String()
}

func (t *Table) interface_() interface{} {
	return t.getv().Interface()
}

//// map op

func (t *Table) mapMap() map[*Table]*Table {
	l := t.getv().Len()
	m := make(map[*Table]*Table, l)
	iter := t.getv().MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		m[&Table{v: k}] = &Table{v: v}
	}
	return m
}

func (t *Table) sliceMap() map[*Table]*Table {
	l := t.getv().Len()
	m := make(map[*Table]*Table, l)
	v := t.getv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		m[&Table{i: i}] = &Table{v: ev}
	}
	return m
}

func (t *Table) structMap() map[*Table]*Table {
	num := t.getv().NumField()
	m := make(map[*Table]*Table, num)
	rv := t.getv()
	rt := rv.Type()
	for i := 0; i < num; i++ {
		fn := rt.Field(i).Name
		fv := rv.Field(i)
		m[&Table{i: fn}] = &Table{v: fv}
	}
	return m
}

//// slice op

func (t *Table) sliceSlice() []*Table {
	l := t.getv().Len()
	s := make([]*Table, l)
	v := t.getv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		s[i] = &Table{v: ev}
	}
	return s
}

func (t *Table) structSlice() []*Table {
	num := t.getv().NumField()
	s := make([]*Table, num)
	rv := t.getv()
	for i := 0; i < num; i++ {
		fv := rv.Field(i)
		s[i] = &Table{v: fv}
	}
	return s
}

//// alist op

func (t *Table) mapAList() [][2]*Table {
	l := t.getv().Len()
	alist := make([][2]*Table, 0, l)
	iter := t.getv().MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		alist = append(alist, [2]*Table{&Table{v: k}, &Table{v: v}})
	}
	return alist
}

func (t *Table) sliceAList() [][2]*Table {
	l := t.getv().Len()
	alist := make([][2]*Table, 0, l)
	v := t.getv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		alist = append(alist, [2]*Table{&Table{i: i}, &Table{v: ev}})
	}
	return alist
}

func (t *Table) structAList() [][2]*Table {
	num := t.getv().NumField()
	alist := make([][2]*Table, 0, num)
	rv := t.getv()
	rt := rv.Type()
	for i := 0; i < num; i++ {
		fn := rt.Field(i).Name
		fv := rv.Field(i)
		alist = append(alist, [2]*Table{&Table{i: fn}, &Table{v: fv}})
	}
	return alist
}

//// plist op

func (t *Table) mapPList() []*Table {
	l := t.getv().Len()
	plist := make([]*Table, 0, 2*l)
	iter := t.getv().MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		plist = append(plist, &Table{v: k}, &Table{v: v})
	}
	return plist
}

func (t *Table) slicePList() []*Table {
	l := t.getv().Len()
	plist := make([]*Table, 0, 2*l)
	v := t.getv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		plist = append(plist, &Table{i: i}, &Table{v: ev})
	}
	return plist
}

func (t *Table) structPList() []*Table {
	num := t.getv().NumField()
	plist := make([]*Table, 0, 2*num)
	rv := t.getv()
	rt := rv.Type()
	for i := 0; i < num; i++ {
		fn := rt.Field(i).Name
		fv := rv.Field(i)
		plist = append(plist, &Table{i: fn}, &Table{v: fv})
	}
	return plist
}
