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

// getiv get indirect value of t
func (t *Table) getiv() reflect.Value {
	// indirect interface{} or Ptr Value
	return indirect(t.getv())
}

func (t *Table) geti() interface{} {
	if t.i == nil {
		t.i = t.v.Interface()
	}
	return t.i
}

//// get op

func (t *Table) mapGet(k interface{}) *Table {
	v := t.getiv().MapIndex(reflect.ValueOf(k))
	if v.Kind() == reflect.Invalid {
		return nil
	}
	return &Table{v: v}
}

func (t *Table) sliceGet(idx int) *Table {
	l := t.getiv().Len()
	if idx >= l {
		return nil
	}

	v := t.getiv().Index(idx)
	return &Table{v: v}
}

func (t *Table) structGet(field string) *Table {
	v := t.getiv().FieldByName(field)
	if v.Kind() == reflect.Invalid {
		return nil
	}
	return &Table{v: v}
}

//// put op

func (t *Table) mapPut(k, v interface{}) error {
	if t.getiv().IsNil() {
		t.v = reflect.MakeMap(t.getiv().Type())
		t.i = t.v.Interface()
	}
	t.getiv().SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	return nil
}

func (t *Table) arrayPut(idx int, v interface{}) error {
	cap := t.getiv().Cap()
	if idx >= cap {
		return nil
	}
	ev := t.getiv().Index(idx)
	ev.Set(reflect.ValueOf(v))
	return nil
}

func (t *Table) slicePut(idx int, v interface{}) error {
	l := t.getiv().Len()
	if idx < l { // set
		ev := t.getiv().Index(idx)
		ev.Set(reflect.ValueOf(v))
	} else { // append
		sv := t.getiv()
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
	fv := t.getiv().FieldByName(fn)
	if fv.IsValid() {
		return &ErrNotExist{"Table.structPut", fn + " field"}
	}
	fv.Set(reflect.ValueOf(v))
	return nil
}

func (t *Table) bool() bool {
	return t.getiv().Bool()
}

func (t *Table) int() int64 {
	return t.getiv().Int()
}

func (t *Table) uint() uint64 {
	return t.getiv().Uint()
}

func (t *Table) float() float64 {
	return t.getiv().Float()
}

func (t *Table) complex_() complex128 {
	return t.getiv().Complex()
}

func (t *Table) string() string {
	return t.getiv().String()
}

func (t *Table) interface_() interface{} {
	return t.getiv().Interface()
}

//// map op

func (t *Table) mapMap() map[*Table]*Table {
	l := t.getiv().Len()
	m := make(map[*Table]*Table, l)
	iter := t.getiv().MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		m[&Table{v: k}] = &Table{v: v}
	}
	return m
}

func (t *Table) sliceMap() map[*Table]*Table {
	l := t.getiv().Len()
	m := make(map[*Table]*Table, l)
	v := t.getiv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		m[&Table{i: i}] = &Table{v: ev}
	}
	return m
}

func (t *Table) structMap() map[*Table]*Table {
	num := t.getiv().NumField()
	m := make(map[*Table]*Table, num)
	rv := t.getiv()
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
	l := t.getiv().Len()
	s := make([]*Table, l)
	v := t.getiv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		s[i] = &Table{v: ev}
	}
	return s
}

func (t *Table) structSlice() []*Table {
	num := t.getiv().NumField()
	s := make([]*Table, num)
	rv := t.getiv()
	for i := 0; i < num; i++ {
		fv := rv.Field(i)
		s[i] = &Table{v: fv}
	}
	return s
}

//// alist op

func (t *Table) mapAList() [][2]*Table {
	l := t.getiv().Len()
	alist := make([][2]*Table, 0, l)
	iter := t.getiv().MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		alist = append(alist, [2]*Table{&Table{v: k}, &Table{v: v}})
	}
	return alist
}

func (t *Table) sliceAList() [][2]*Table {
	l := t.getiv().Len()
	alist := make([][2]*Table, 0, l)
	v := t.getiv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		alist = append(alist, [2]*Table{&Table{i: i}, &Table{v: ev}})
	}
	return alist
}

func (t *Table) structAList() [][2]*Table {
	num := t.getiv().NumField()
	alist := make([][2]*Table, 0, num)
	rv := t.getiv()
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
	l := t.getiv().Len()
	plist := make([]*Table, 0, 2*l)
	iter := t.getiv().MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		plist = append(plist, &Table{v: k}, &Table{v: v})
	}
	return plist
}

func (t *Table) slicePList() []*Table {
	l := t.getiv().Len()
	plist := make([]*Table, 0, 2*l)
	v := t.getiv()
	for i := 0; i < l; i++ {
		ev := v.Index(i)
		plist = append(plist, &Table{i: i}, &Table{v: ev})
	}
	return plist
}

func (t *Table) structPList() []*Table {
	num := t.getiv().NumField()
	plist := make([]*Table, 0, 2*num)
	rv := t.getiv()
	rt := rv.Type()
	for i := 0; i < num; i++ {
		fn := rt.Field(i).Name
		fv := rv.Field(i)
		plist = append(plist, &Table{i: fn}, &Table{v: fv})
	}
	return plist
}
