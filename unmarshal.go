package table

import (
	"math/bits"
	"reflect"
)

var (
	intLevel = map[reflect.Kind]int{
		reflect.Int8:  1,
		reflect.Int16: 2,
		reflect.Int32: 3,
		reflect.Int64: 4,
	}
	uintLevel = map[reflect.Kind]int{
		reflect.Uint8:  1,
		reflect.Uint16: 2,
		reflect.Uint32: 3,
		reflect.Uint64: 4,
	}
	floatLevel = map[reflect.Kind]int{
		reflect.Float32: 1,
		reflect.Float64: 2,
	}
	complexLevel = map[reflect.Kind]int{
		reflect.Complex64:  1,
		reflect.Complex128: 2,
	}
)

// Unmarshal ...
func (t *Table) Unmarshal(value interface{}) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return &ErrUnsupportedKind{"Table.Unmarshal", v.Kind()}
	}
	v = v.Elem()
	return t.unmarshal(v)
}

func (t *Table) unmarshal(v reflect.Value) (err error) {
	switch v.Kind() {
	case reflect.Bool:
		return t.unmarshalBool(v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return t.unmarshalInt(v)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return t.unmarshalUint(v)

	case reflect.Float32, reflect.Float64:
		return t.unmarshalFloat(v)

	case reflect.Complex64, reflect.Complex128:
		return t.unmarshalComplex(v)

	case reflect.String:
		return t.unmarshalString(v)

	case reflect.Map:
		return t.unmarshalMap(v)

	case reflect.Array:
		return t.unmarshalArray(v)

	case reflect.Slice:
		return t.unmarshalSlice(v)

	case reflect.Struct:
		return t.unmarshalStruct(v)

	case reflect.Interface:
		return t.unmarshalInterface(v)

	default:
		return &ErrUnsupportedKind{"Table.unmarshal", v.Kind()}
	}
}

func (t *Table) unmarshalInterface(v reflect.Value) error {
	v.Set(t.getv())
	return nil
}

func (t *Table) unmarshalBool(v reflect.Value) error {
	v.SetBool(t.bool())
	return nil
}

func (t *Table) unmarshalInt(v reflect.Value) error {
	tk := t.getv().Kind()
	vk0 := v.Kind()
	vk := reflect.Invalid

	if vk0 == reflect.Int { // convert Int to Int32 or Int64
		if bits.UintSize == 32 {
			vk = reflect.Int32
		}
		if bits.UintSize == 64 {
			vk = reflect.Int64
		}
	} else {
		vk = vk0
	}

	if intLevel[vk] < intLevel[tk] {
		return &ErrTypeUnequal{
			"Table.unmarshalInt",
			vk,
			tk,
		}
	}

	v.SetInt(t.int())
	return nil
}

func (t *Table) unmarshalUint(v reflect.Value) error {
	tk := t.getv().Kind()
	vk0 := v.Kind()
	vk := reflect.Invalid

	if vk0 == reflect.Uint { // convert Uint to Uint32 or Uint64
		if bits.UintSize == 32 {
			vk = reflect.Uint32
		}
		if bits.UintSize == 64 {
			vk = reflect.Uint64
		}
	} else {
		vk = vk0
	}

	if uintLevel[vk] < uintLevel[tk] {
		return &ErrTypeUnequal{
			"Table.unmarshalUint",
			vk,
			tk,
		}
	}

	v.SetUint(t.uint())
	return nil
}

func (t *Table) unmarshalFloat(v reflect.Value) error {
	tk := t.getv().Kind()
	vk := v.Kind()

	if floatLevel[vk] < floatLevel[tk] {
		return &ErrTypeUnequal{
			"Table.unmarshalFloat",
			tk,
			vk,
		}
	}

	v.SetFloat(t.float())
	return nil
}

func (t *Table) unmarshalComplex(v reflect.Value) error {
	tk := t.getv().Kind()
	vk := v.Kind()

	if complexLevel[vk] < complexLevel[tk] {
		return &ErrTypeUnequal{
			"Table.unmarshalComplex",
			tk,
			vk,
		}
	}

	v.SetComplex(t.complex_())
	return nil
}

func (t *Table) unmarshalString(v reflect.Value) error {
	v.SetString(t.string())
	return nil
}

func (t *Table) unmarshalMap(m reflect.Value) error {
	tm, err := t.Map()
	if err != nil {
		return err
	}

	for k, v := range tm {
		mk := k.getv()
		mv := m.MapIndex(mk)
		if mv.Kind() == reflect.Invalid {
			mv = reflect.New(m.Type().Elem())
			mv = mv.Elem()
		}

		if err := v.unmarshal(mv); err != nil {
			return err
		}
		m.SetMapIndex(mk, mv)
	}
	return nil
}

func (t *Table) unmarshalArray(a reflect.Value) error {
	ts, err := t.Slice()
	if err != nil {
		return err
	}

	for i, v := range ts {
		if i >= a.Cap() {
			break
		}

		ev := a.Index(i)

		if err := v.unmarshal(ev); err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) unmarshalSlice(s reflect.Value) error {
	ts, err := t.Slice()
	if err != nil {
		return err
	}

	// l := s.Len()
	newSlice := reflect.MakeSlice(s.Type(), len(ts), s.Len()+len(ts))
	for i, v := range ts {
		var ev reflect.Value
		if i < newSlice.Len() {
			ev = newSlice.Index(i)
		} else {
			ev = reflect.New(newSlice.Type().Elem())
			newSlice = reflect.Append(newSlice, ev)
			ev = newSlice.Index(i)
		}

		if err := v.unmarshal(ev); err != nil {
			return err
		}
	}
	s.Set(newSlice)
	return nil
}
func (t *Table) unmarshalStruct(s reflect.Value) error {
	tm, err := t.Map()
	if err != nil {
		return err
	}

	stype := s.Type()
	tag2Fname := map[string]string{}
	passedFnames := map[string]bool{}
	for i := 0; i < stype.NumField(); i++ {
		sf := stype.Field(i)
		tag := sf.Tag.Get("table")
		if tag == "_" {
			passedFnames[sf.Name] = true
		} else if tag != "" {
			tag2Fname[tag] = sf.Name
		}
	}
	for k, v := range tm {
		key := k.String()
		if passedFnames[key] == true {
			continue
		}
		fn := tag2Fname[key]
		if fn == "" {
			fn = key
		}
		f := s.FieldByName(fn)
		if f.Kind() == reflect.Invalid {
			continue
		}
		if err := v.unmarshal(f); err != nil {
			return err
		}
		passedFnames[fn] = true
	}
	return nil
}
