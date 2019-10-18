package table

import (
	"math/bits"
	"reflect"
	"time"
)

var (
	// TimeLayout default time layout
	TimeLayout = "Mon Jan 2 15:04:05 -0700 MST 2006"
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

// Conv ...
func (t *Table) Conv(value interface{}) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return &ErrUnsupportedKind{"Table.Conv", v.Kind()}
	}
	v = v.Elem()
	return t.conv(v)
}

func (t *Table) conv(v reflect.Value) (err error) {
	switch v.Type().String() {
	case "time.Duration":
		return t.convTimeDuration(v)
	case "time.Time":
		return t.convTime(v)
	}

	switch v.Kind() {
	case reflect.Bool:
		return t.convBool(v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return t.convInt(v)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return t.convUint(v)

	case reflect.Float32, reflect.Float64:
		return t.convFloat(v)

	case reflect.Complex64, reflect.Complex128:
		return t.convComplex(v)

	case reflect.String:
		return t.convString(v)

	case reflect.Map:
		return t.convMap(v)

	case reflect.Array:
		return t.convArray(v)

	case reflect.Slice:
		return t.convSlice(v)

	case reflect.Struct:
		return t.convStruct(v)

	case reflect.Interface:
		return t.convInterface(v)

	case reflect.Ptr:
		return t.convPtr(v)

	default:
		return &ErrUnsupportedKind{"Table.conv", v.Kind()}
	}
}

func (t *Table) convTimeDuration(v reflect.Value) error {
	td, err := time.ParseDuration(t.String())
	if err != nil {
		return err
	}
	v.SetInt(int64(td))
	return nil
}

func (t *Table) convTime(v reflect.Value) error {
	timex, err := time.Parse(TimeLayout, t.String())
	if err != nil {
		return err
	}
	v.Set(reflect.ValueOf(timex))
	return nil
}

func (t *Table) convPtr(v reflect.Value) error {
	rv := v
	if v.IsNil() {
		rv = reflect.New(v.Type().Elem())
	}
	if err := t.conv(rv.Elem()); err != nil {
		return err
	}
	v.Set(rv)

	return nil
}

func (t *Table) convInterface(v reflect.Value) error {
	v.Set(t.getv())
	return nil
}

func (t *Table) convBool(v reflect.Value) error {
	v.SetBool(t.bool())
	return nil
}

func (t *Table) convInt(v reflect.Value) error {
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
			"Table.convInt",
			vk,
			tk,
		}
	}

	v.SetInt(t.int())
	return nil
}

func (t *Table) convUint(v reflect.Value) error {
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
			"Table.convUint",
			vk,
			tk,
		}
	}

	v.SetUint(t.uint())
	return nil
}

func (t *Table) convFloat(v reflect.Value) error {
	tk := t.getv().Kind()
	vk := v.Kind()

	if floatLevel[vk] < floatLevel[tk] {
		return &ErrTypeUnequal{
			"Table.convFloat",
			tk,
			vk,
		}
	}

	v.SetFloat(t.float())
	return nil
}

func (t *Table) convComplex(v reflect.Value) error {
	tk := t.getv().Kind()
	vk := v.Kind()

	if complexLevel[vk] < complexLevel[tk] {
		return &ErrTypeUnequal{
			"Table.convComplex",
			tk,
			vk,
		}
	}

	v.SetComplex(t.complex_())
	return nil
}

func (t *Table) convString(v reflect.Value) error {
	v.SetString(t.string())
	return nil
}

func (t *Table) convMap(m reflect.Value) error {
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

		if err := v.conv(mv); err != nil {
			return err
		}
		m.SetMapIndex(mk, mv)
	}
	return nil
}

func (t *Table) convArray(a reflect.Value) error {
	ts, err := t.Slice()
	if err != nil {
		return err
	}

	for i, v := range ts {
		if i >= a.Cap() {
			break
		}

		ev := a.Index(i)

		if err := v.conv(ev); err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) convSlice(s reflect.Value) error {
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

		if err := v.conv(ev); err != nil {
			return err
		}
	}
	s.Set(newSlice)
	return nil
}
func (t *Table) convStruct(s reflect.Value) error {
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
		if err := v.conv(f); err != nil {
			return err
		}
		passedFnames[fn] = true
	}
	return nil
}
