package table

import (
	"fmt"
	"math/bits"
	"reflect"
)

// Table ...
type Table struct {
	i interface{}
	v reflect.Value
}

// New new a Table from v
func New(v interface{}) *Table {
	return &Table{i: v}
}

// Get returns the value with the given key.
//
// If t's kind is Map, Get returns the value associated with key in the map.
// If t's kind is Array or Slice, Get returns t's k'th element, the k must be int.
// If t's kind is Struct, Get returns the struct field with the given field name, the k must be string.
// if t's kind is Interface or Ptr, indirect it.
// It returns the nil if k is not found in the t.
// It returns error if t's kind is not Map, Array, Slice or Struct.
func (t *Table) Get(k interface{}) (*Table, error) {
	v := t.getv()
	switch v.Kind() {
	case reflect.Map:
		return t.mapGet(k), nil
	case reflect.Array, reflect.Slice:
		return t.sliceGet(k.(int)), nil
	case reflect.Struct:
		return t.structGet(k.(string)), nil
	case reflect.Interface, reflect.Ptr:
		vt := &Table{v: indirect(v)}
		return vt.Get(k)
	default:
		return nil, &ErrUnsupportedKind{"Table.Get", v.Kind()}
	}
}

// Set set t's value to v.
//
// If t's value can't setable, returns ErrCannotSet.
// If t's kind and v's kind is not equivalence, returns ErrTypeUnequal.
// It returns nil, that set successful.
//
// If set map key, struct field or array/slice index, using Table.Put.
//
//  TODO: balala
func (t *Table) Set(v interface{}) error {
	tv := t.getv()
	if tv.Kind() == reflect.Interface || tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	if !tv.CanSet() {
		return &ErrCannotSet{"Table.Set"}
	}

	vv := reflect.ValueOf(v)
	if tv.Kind() != vv.Kind() {
		return &ErrTypeUnequal{"Table.Set", tv.Kind(), vv.Kind()}
	}

	tv.Set(vv)

	// reset
	t.i = nil

	return nil
}

// Put put k, v to map, array, slice or struct(structed type).
//
// If t's kind is map, the k indicates key of map.
// If t's kind is array/slice, the k indecates index of array/slice.
// If t's kind is struct, the k indecates fieldname of struct.
//
// If k in t, and set k's value to v.
//
// If t's kind is not map, array, slice or struct, returns ErrUnsupportedKind.
func (t *Table) Put(k, v interface{}) (err error) {
	switch t.getv().Kind() {
	case reflect.Map:
		return t.mapPut(k, v)
	case reflect.Array:
		return t.arrayPut(k.(int), v)
	case reflect.Slice:
		return t.slicePut(k.(int), v)
	case reflect.Struct:
		return t.structPut(k.(string), v)
	default:
		return &ErrUnsupportedKind{"Table.Put", t.getv().Kind()}
	}
}

// TODO
// Bytes returns t's underlying value as a []bytes.
// It returns error if t's underlying value is not a slice of bytes.
func (t *Table) Bytes() ([]byte, error) {
	if t.getv().Kind() != reflect.Slice {
		return nil, &ErrUnsupportedKind{"Table.Bytes", t.getv().Kind()}
	}
	v := t.getv()
	elemk := v.Type().Elem().Kind()
	if elemk != reflect.Uint8 {
		return nil, &ErrUnsupportedKind{"Table.Bytes", "slice of " + elemk.String()}
	}
	return v.Bytes(), nil
}

// Bool returns t's underlying value.
// It returns error if t's kind is not Bool.
func (t *Table) Bool() (bool, error) {
	switch t.getv().Kind() {
	case reflect.Bool:
		return t.bool(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Bool()
	default:
		return false, &ErrUnsupportedKind{"Table.Int", t.getv().Kind()}
	}
}

// Int returns t's underlying value as an int.
// It returns error if t's kind is not Int, Int8, Int16, Int32, Uint8 or Uint16,
// and if t's kind is Int64 or Uint32 also Int is 32 bits.
func (t *Table) Int() (i int, err error) {
	switch t.getv().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		i = int(t.int())
	case reflect.Int64:
		if bits.UintSize == 64 { // if int size is 64
			i = int(t.int())
		} else { // if int size is 32
			err = &ErrUnsupportedKind{"Table.Int", t.getv().Kind()}
		}

	case reflect.Uint8, reflect.Uint16:
		i = int(t.uint())

	case reflect.Uint32:
		if bits.UintSize == 64 {
			i = int(t.uint())
		} else {
			err = &ErrUnsupportedKind{"Table.Int", t.getv().Kind()}
		}

	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Int()

	default:
		err = &ErrUnsupportedKind{"Table.Int", t.getv().Kind()}
	}
	return
}

// Int8 returns t's underlying value as an int8.
// It returns error if t's kind is not Int8.
func (t *Table) Int8() (int8, error) {
	switch t.getv().Kind() {
	case reflect.Int8:
		return int8(t.int()), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Int8()
	default:
		return 0, &ErrUnsupportedKind{"Table.Int8", t.getv().Kind()}
	}
}

// Int16 returns t's underlying value as an int16.
// It returns error if t's kind is not Int, Int8, Int16, or Uint8.
func (t *Table) Int16() (int16, error) {
	switch t.getv().Kind() {
	case reflect.Int8, reflect.Int16:
		return int16(t.int()), nil
	case reflect.Uint8:
		return int16(t.uint()), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Int16()
	default:
		return 0, &ErrUnsupportedKind{"Table.Int16", t.getv().Kind()}
	}
}

// Int32 returns t's underlying value as an int32.
// It returns error if t's kind is not Int, Int8, Int16, Int32, Uint8 or Uint16,
// and if t's kind is Int also Int is 64 bits.
func (t *Table) Int32() (int32, error) {
	switch t.getv().Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32:
		return int32(t.int()), nil
	case reflect.Int:
		if bits.UintSize == 32 { // 32
			return int32(t.int()), nil
		}
		return 0, &ErrUnsupportedKind{"Table.Int32", t.getv().Kind()}

	case reflect.Uint8, reflect.Uint16:
		return int32(t.uint()), nil

	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Int32()

	default:
		return 0, &ErrUnsupportedKind{"Table.Int32", t.getv().Kind()}
	}
}

// Int64 returns t's underlying value as an int64.
// It returns error if t's kind is not Int, Int8, Int16, Int32, Uint8, Uint16, Uint32
// and if t's kind is Uint also Uint is 64 bits.
func (t *Table) Int64() (int64, error) {
	switch t.getv().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return t.int(), nil

	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int64(t.uint()), nil
	case reflect.Uint:
		if bits.UintSize == 32 { // 32
			return int64(t.uint()), nil
		}
		return 0, &ErrUnsupportedKind{"Table.Int64", t.getv().Kind()}

	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Int64()

	default:
		return 0, &ErrUnsupportedKind{"Table.Int64", t.getv().Kind()}
	}
}

// Uint returns t's underlying value as an uint.
// It returns error if t's kind is not Uint, Uint8, Uint16 or Uint32,
// and if t's kind is Uint64 also Uint is 32 bits.
func (t *Table) Uint() (i uint, err error) {
	switch t.getv().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		i = uint(t.uint())
	case reflect.Uint64:
		if bits.UintSize == 64 { // 64
			i = uint(t.uint())
		} else { // 32
			err = &ErrUnsupportedKind{"Table.Uint", t.getv().Kind()}
		}
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Uint()

	default:
		err = &ErrUnsupportedKind{"Table.Uint", t.getv().Kind()}
	}
	return
}

// Uint8 returns t's underlying value as an uint8.
// It returns error if t's kind is not Uint8.
func (t *Table) Uint8() (uint8, error) {
	switch t.getv().Kind() {
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Uint8()
	case reflect.Uint8:
		return uint8(t.uint()), nil
	default:
		return 0, &ErrUnsupportedKind{"Table.Uint8", t.getv().Kind()}
	}
}

// Uint16 returns t's underlying value as an uint16.
// It returns error if t's kind is not Uint8 or Uint16.
func (t *Table) Uint16() (uint16, error) {
	switch t.getv().Kind() {
	case reflect.Uint8, reflect.Uint16:
		return uint16(t.uint()), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Uint16()
	default:
		return 0, &ErrUnsupportedKind{"Table.Uint16", t.getv().Kind()}
	}
}

// Uint32 returns t's underlying value as an uint32.
// It returns error if t's kind is not Uint8, Uint16 or Uint32,
// and if t's kind is Uint also Uint is 64 bits.
func (t *Table) Uint32() (uint32, error) {
	switch t.getv().Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return uint32(t.uint()), nil

	case reflect.Uint:
		if bits.UintSize == 32 { // 32
			return uint32(t.uint()), nil
		}
		return 0, &ErrUnsupportedKind{"Table.Uint32", t.getv().Kind()}

	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Uint32()

	default:
		return 0, &ErrUnsupportedKind{"Table.Uint32", t.getv().Kind()}
	}
}

// Uint64 returns t's underlying value as an uint64.
// It returns error if t's kind is not Uint*.
func (t *Table) Uint64() (uint64, error) {
	switch t.getv().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return t.uint(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Uint64()
	default:
		return 0, &ErrUnsupportedKind{"Table.Uint64", t.getv().Kind()}
	}
}

// Float32 returns t's underlying value as an float32.
// It returns error if t's kind is not Uint*, Int* or Float32.
func (t *Table) Float32() (float32, error) {
	switch t.getv().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float32(t.int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float32(t.uint()), nil
	case reflect.Float32:
		return float32(t.float()), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Float32()
	default:
		return 0, &ErrUnsupportedKind{"Table.Float32", t.getv().Kind()}
	}
}

// Float64 returns t's underlying value as an float64.
// It returns error if t's kind is not Uint*, Int* or Float*.
func (t *Table) Float64() (float64, error) {
	switch t.getv().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(t.int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(t.uint()), nil
	case reflect.Float32, reflect.Float64:
		return t.float(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Float64()
	default:
		return 0, &ErrUnsupportedKind{"Table.Float64", t.getv().Kind()}
	}
}

// Complex64 returns t's underlying value as an complex64.
// It returns error if t's kind is not Uint*, Int*, Float32 or Complex64.
func (t *Table) Complex64() (complex64, error) {
	switch t.getv().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return complex(float32(t.uint()), 0), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex(float32(t.int()), 0), nil
	case reflect.Float32:
		return complex(float32(t.float()), 0), nil
	case reflect.Complex64:
		return complex64(t.complex_()), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Complex64()
	default:
		return 0i, &ErrUnsupportedKind{"Table.Complex64", t.getv().Kind()}
	}
}

// Complex128 returns t's underlying value as an complex128.
// It returns error if t's kind is not Uint*, Int*, Float* or Complex*.
func (t *Table) Complex128() (complex128, error) {
	switch t.getv().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return complex(float64(t.uint()), 0), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex(float64(t.int()), 0), nil
	case reflect.Float32, reflect.Float64:
		return complex(t.float(), 0), nil
	case reflect.Complex64, reflect.Complex128:
		return t.complex_(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Complex128()
	default:
		return 0i, &ErrUnsupportedKind{"Table.Complex128", t.getv().Kind()}
	}
}

// Map returns t's underlying value as a map.
// It returns error if t's kind is not Map, Array, Slice or Struct.
func (t *Table) Map() (map[*Table]*Table, error) {
	switch t.getv().Kind() {
	case reflect.Map:
		return t.mapMap(), nil
	case reflect.Array, reflect.Slice:
		return t.sliceMap(), nil
	case reflect.Struct:
		return t.structMap(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Map()
	default:
		return nil, &ErrUnsupportedKind{"Table.Map", t.getv().Kind()}
	}
}

// Slice returns t's underlying value as a slice.
// It returns error if t's kind is not Array, Slice or Struct.
func (t *Table) Slice() ([]*Table, error) {
	switch t.getv().Kind() {
	case reflect.Array, reflect.Slice:
		return t.sliceSlice(), nil
	case reflect.Struct:
		return t.structSlice(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).Slice()
	default:
		return nil, &ErrUnsupportedKind{"Table.Slice", t.getv().Kind()}
	}
}

// AList returns t's underlying value as an association list.
// It returns error if t's kind is not Map, Array, Slice or Struct.
func (t *Table) AList() ([][2]*Table, error) {
	switch t.getv().Kind() {
	case reflect.Map:
		return t.mapAList(), nil
	case reflect.Array, reflect.Slice:
		return t.sliceAList(), nil
	case reflect.Struct:
		return t.structAList(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).AList()
	default:
		return nil, &ErrUnsupportedKind{"Table.AList", t.getv().Kind()}
	}
}

// PList returns t's underlying value as an property list.
// It returns error if t's kind is not Map, Array, Slice or Struct.
func (t *Table) PList() ([]*Table, error) {
	switch t.getv().Kind() {
	case reflect.Map:
		return t.mapPList(), nil
	case reflect.Array, reflect.Slice:
		return t.slicePList(), nil
	case reflect.Struct:
		return t.structPList(), nil
	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).PList()
	default:
		return nil, &ErrUnsupportedKind{"Table.PList", t.getv().Kind()}
	}
}

func (t *Table) Interface() interface{} {
	return t.geti()
}

func (t *Table) Ptr() uintptr {
	return t.getv().Pointer()
}

func (t *Table) String() string {
	switch t.getv().Kind() {
	case reflect.Invalid:
		return ""
	case reflect.String:
		return t.string()
	case reflect.Bool:
		return fmt.Sprintf("%t", t.bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", t.int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%d", t.uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", t.float())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%g", t.complex_())
	case reflect.Chan:
		return fmt.Sprintf("%p", t.getv().Interface())
	case reflect.UnsafePointer:
		return fmt.Sprintf("%x", t.getv().Interface())
	case reflect.Struct:
		return fmt.Sprintf("%#v", t.getv().Interface())
	case reflect.Slice, reflect.Array, reflect.Map:
		return fmt.Sprintf("%v", t.getv().Interface())
	default:
		return t.string()
	}
}

type eachDoFunc func(k, v *Table) error

func (t *Table) EachDo(f eachDoFunc) error {
	switch t.getv().Kind() {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		m, err := t.Map()
		if err != nil {
			return err
		}
		for k, v := range m {
			if err := f(k, v); err != nil {
				return err
			}
		}
	case reflect.Chan:
		idx := 0
		for {
			v, ok := t.getv().Recv()
			if !ok {
				break
			}

			if err := f(&Table{i: idx}, &Table{v: v}); err != nil {
				return err
			}

			idx++
		}
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:

		if err := f(nil, t); err != nil {
			return err
		}

	case reflect.Interface, reflect.Ptr:
		return (&Table{v: indirect(t.getv())}).EachDo(f)

	default:
		return &ErrUnsupportedKind{"Table.EachDo", t.getv().Kind()}
	}
	return nil
}
