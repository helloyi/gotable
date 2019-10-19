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

// ConvTo convToert t to value
func (t *Table) ConvTo(value interface{}) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return &ErrUnsupportedKind{"Table.ConvTo", v.Kind()}
	}
	v = v.Elem()
	return t.convTo(v)
}

func (t *Table) convTo(v reflect.Value) (err error) {
	switch v.Type().String() {
	case "time.Duration":
		return t.convToTimeDuration(v)
	case "time.Time":
		return t.convToTime(v)
	}

	switch v.Kind() {
	case reflect.Bool:
		return t.convToBool(v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return t.convToInt(v)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return t.convToUint(v)

	case reflect.Float32, reflect.Float64:
		return t.convToFloat(v)

	case reflect.Complex64, reflect.Complex128:
		return t.convToComplex(v)

	case reflect.String:
		return t.convToString(v)

	case reflect.Map:
		return t.convToMap(v)

	case reflect.Array:
		return t.convToArray(v)

	case reflect.Slice:
		return t.convToSlice(v)

	case reflect.Struct:
		return t.convToStruct(v)

	case reflect.Interface:
		return t.convToInterface(v)

	case reflect.Ptr:
		return t.convToPtr(v)

	default:
		return &ErrUnsupportedKind{"Table.convTo", v.Kind()}
	}
}

func (t *Table) convToTimeDuration(v reflect.Value) error {
	s, err := t.String()
	if err != nil {
		return err
	}
	td, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	v.SetInt(int64(td))
	return nil
}

func (t *Table) convToTime(v reflect.Value) error {
	s, err := t.String()
	if err != nil {
		return err
	}
	timex, err := time.Parse(TimeLayout, s)
	if err != nil {
		return err
	}
	v.Set(reflect.ValueOf(timex))
	return nil
}

func (t *Table) convToPtr(v reflect.Value) error {
	rv := v
	if v.IsNil() {
		rv = reflect.New(v.Type().Elem())
	}
	if err := t.convTo(rv.Elem()); err != nil {
		return err
	}
	v.Set(rv)

	return nil
}

func (t *Table) convToInterface(v reflect.Value) error {
	v.Set(t.getv())
	return nil
}

func (t *Table) convToBool(v reflect.Value) error {
	b, err := t.Bool()
	if err != nil {
		return err
	}

	v.SetBool(b)
	return nil
}

func (t *Table) convToInt(v reflect.Value) error {
	tk := t.getv().Kind()
	vk0 := v.Kind()
	vk := reflect.Invalid

	if vk0 == reflect.Int { // convToert Int to Int32 or Int64
		if bits.UintSize == 32 {
			vk = reflect.Int32
		}
		if bits.UintSize == 64 {
			vk = reflect.Int64
		}
	} else {
		vk = vk0
	}

	iv, err := t.Int64()
	if err != nil {
		return err
	}
	if intLevel[vk] < intLevel[tk] {
		return &ErrTypeUnequal{
			"Table.convToInt",
			vk,
			tk,
		}
	}

	v.SetInt(iv)
	return nil
}

func (t *Table) convToUint(v reflect.Value) error {
	tk := t.getv().Kind()
	vk0 := v.Kind()
	vk := reflect.Invalid

	if vk0 == reflect.Uint { // convToert Uint to Uint32 or Uint64
		if bits.UintSize == 32 {
			vk = reflect.Uint32
		}
		if bits.UintSize == 64 {
			vk = reflect.Uint64
		}
	} else {
		vk = vk0
	}

	uv, err := t.Uint64()
	if err != nil {
		return err
	}
	if uintLevel[vk] < uintLevel[tk] {
		return &ErrTypeUnequal{
			"Table.convToUint",
			vk,
			tk,
		}
	}

	v.SetUint(uv)
	return nil
}

func (t *Table) convToFloat(v reflect.Value) error {
	tk := t.getv().Kind()
	vk := v.Kind()

	tf, err := t.Float64()
	if err != nil {
		return err
	}
	if floatLevel[vk] < floatLevel[tk] {
		return &ErrTypeUnequal{
			"Table.convToFloat",
			tk,
			vk,
		}
	}

	v.SetFloat(tf)
	return nil
}

func (t *Table) convToComplex(v reflect.Value) error {
	tk := t.getv().Kind()
	vk := v.Kind()

	tc, err := t.Complex128()
	if err != nil {
		return err
	}
	if complexLevel[vk] < complexLevel[tk] {
		return &ErrTypeUnequal{
			"Table.convToComplex",
			tk,
			vk,
		}
	}

	v.SetComplex(tc)
	return nil
}

func (t *Table) convToString(v reflect.Value) error {
	ts, err := t.String()
	if err != nil {
		return err
	}
	v.SetString(ts)
	return nil
}

func (t *Table) convToMap(m reflect.Value) error {
	tm, err := t.Map()
	if err != nil {
		return err
	}

	if m.IsNil() {
		m.Set(reflect.MakeMap(m.Type()))
	}
	for k, v := range tm {
		mk := k.getv()
		mv := m.MapIndex(mk)
		if mv.Kind() == reflect.Invalid {
			mv = reflect.New(m.Type().Elem())
			mv = mv.Elem()
		}

		if err := v.convTo(mv); err != nil {
			return err
		}
		m.SetMapIndex(mk, mv)
	}
	return nil
}

func (t *Table) convToArray(a reflect.Value) error {
	ts, err := t.Slice()
	if err != nil {
		return err
	}

	for i, v := range ts {
		if i >= a.Cap() {
			break
		}

		ev := a.Index(i)

		if err := v.convTo(ev); err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) convToSlice(s reflect.Value) error {
	ts, err := t.Slice()
	if err != nil {
		return err
	}

	newSlice := reflect.MakeSlice(s.Type(), len(ts), s.Len()+len(ts))
	for i := 0; i < s.Len(); i++ {
		newSlice = reflect.Append(newSlice, s.Index(i))
	}
	for i, v := range ts {
		var ev reflect.Value
		if i < newSlice.Len() {
			ev = newSlice.Index(i)
		} else {
			ev = reflect.New(newSlice.Type().Elem())
			newSlice = reflect.Append(newSlice, ev)
			ev = newSlice.Index(i)
		}

		if err := v.convTo(ev); err != nil {
			return err
		}
	}
	s.Set(newSlice)
	return nil
}
func (t *Table) convToStruct(s reflect.Value) error {
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
		key, err := k.String()
		if err != nil {
			return err
		}
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
		if err := v.convTo(f); err != nil {
			return err
		}
		passedFnames[fn] = true
	}
	return nil
}
