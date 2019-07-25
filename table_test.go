package table

import (
	"math/bits"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ExpectErr(rets ...interface{}) Assertion {
	return Expect(rets[1])
}

var _ = Describe("Table", func() {
	When("bool type", func() {
		Specify("Bool", func() {
			b := true
			t := New(b)
			Expect(t.Bool()).To(Equal(b))

			// other type
			st := New("test")
			ExpectErr(st.Bool()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	When("uint type", func() {
		Specify("Uint", func() {
			x := uint(12)
			ts := []*Table{
				New(x),
				New(uint8(x)),
				New(uint16(x)),
				New(uint32(x)),
			}
			for _, t := range ts {
				Expect(t.Uint()).To(Equal(x))
			}

			ut64 := New(uint64(x))
			if bits.UintSize == 64 {
				Expect(ut64.Uint()).To(Equal(x))
			} else {
				ExpectErr(ut64.Uint()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
			}

			// other type
			st := New("test")
			ExpectErr(st.Uint()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Uint8", func() {
			x := uint8(12)
			ts := []*Table{
				New(uint8(x)),
			}
			for _, t := range ts {
				Expect(t.Uint8()).To(Equal(x))
			}

			// other type
			st := New("test")
			ExpectErr(st.Uint8()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Uint16", func() {
			x := uint16(12)
			ts := []*Table{
				New(uint8(x)),
				New(uint16(x)),
			}
			for _, t := range ts {
				Expect(t.Uint16()).To(Equal(x))
			}

			// other type
			st := New("test")
			ExpectErr(st.Uint16()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Uint32", func() {
			x := uint32(12)
			ts := []*Table{
				New(uint8(x)),
				New(uint16(x)),
				New(uint32(x)),
			}
			for _, t := range ts {
				Expect(t.Uint32()).To(Equal(x))
			}

			ut := New(uint(x))
			if bits.UintSize == 32 {
				Expect(ut.Uint32()).To(Equal(x))
			} else {
				ExpectErr(ut.Uint32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
			}

			// other type
			st := New("test")
			ExpectErr(st.Uint32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Uint64", func() {
			x := uint64(12)
			ts := []*Table{
				New(uint(x)),
				New(uint8(x)),
				New(uint16(x)),
				New(uint32(x)),
				New(uint64(x)),
			}
			for _, t := range ts {
				Expect(t.Uint64()).To(Equal(x))
			}

			// other type
			st := New("test")
			ExpectErr(st.Uint64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	When("int type", func() {
		Specify("Int", func() {
			x := 12
			ts := []*Table{
				New(int(x)),
				New(int8(x)),
				New(int16(x)),
				New(int32(x)),
				New(uint8(x)),
				New(uint16(x)),
			}
			for _, t := range ts {
				Expect(t.Int()).To(Equal(x))
			}

			t64 := New(int64(x))
			if bits.UintSize == 64 {
				Expect(t64.Int()).To(Equal(x))
			} else {
				ExpectErr(t64.Int()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
			}

			ut32 := New(uint32(x))
			if bits.UintSize == 64 {
				Expect(ut32.Int()).To(Equal(x))
			} else {
				ExpectErr(ut32.Int()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
			}

			// other type
			st := New("test")
			ExpectErr(st.Int()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Int8", func() {
			t8 := New(int8(12))
			Expect(t8.Int8()).To(Equal(int8(12)))

			// other type
			ts := New("test")
			ExpectErr(ts.Int8()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Int16", func() {
			x := int16(12)
			ts := []*Table{
				New(int8(x)),
				New(int16(x)),
				New(uint8(x)),
			}
			for _, t := range ts {
				Expect(t.Int16()).To(Equal(x))
			}

			// other type
			st := New("test")
			ExpectErr(st.Int16()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Int32", func() {
			x := int32(12)
			ts := []*Table{
				New(int8(x)),
				New(int16(x)),
				New(int32(x)),
				New(uint8(x)),
				New(uint16(x)),
			}
			for _, t := range ts {
				Expect(t.Int32()).To(Equal(x))
			}

			t := New(int(x))
			if bits.UintSize == 32 {
				Expect(t.Int32()).To(Equal(x))
			} else {
				ExpectErr(t.Int32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
			}

			// other type
			st := New("test")
			ExpectErr(st.Int32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Int64", func() {
			x := int64(12)
			ts := []*Table{
				New(int(x)),
				New(int8(x)),
				New(int16(x)),
				New(int32(x)),
				New(int64(x)),
				New(uint8(x)),
				New(uint16(x)),
				New(uint32(x)),
			}
			for _, t := range ts {
				Expect(t.Int64()).To(Equal(x))
			}

			ut := New(uint(x))
			if bits.UintSize == 32 {
				Expect(ut.Int64()).To(Equal(x))
			} else {
				ExpectErr(ut.Int64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
			}

			// other type
			st := New("test")
			ExpectErr(st.Int64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	When("float type", func() {
		Specify("Float32", func() {
			x := float32(12)
			ts := []*Table{
				New(x),
				New(int(x)),
				New(int8(x)),
				New(int16(x)),
				New(int32(x)),
				New(int64(x)),
				New(uint(x)),
				New(uint8(x)),
				New(uint16(x)),
				New(uint32(x)),
				New(uint64(x)),
			}
			for _, t := range ts {
				Expect(t.Float32()).To(Equal(x))
			}

			// other type
			st := New("test")
			ExpectErr(st.Float32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Float64", func() {
			x := float64(12)
			ts := []*Table{
				New(x),
				New(float32(x)),
				New(int(x)),
				New(int8(x)),
				New(int16(x)),
				New(int32(x)),
				New(int64(x)),
				New(uint(x)),
				New(uint8(x)),
				New(uint16(x)),
				New(uint32(x)),
				New(uint64(x)),
			}
			for _, t := range ts {
				Expect(t.Float64()).To(Equal(x))
			}

			// other type
			st := New("test")
			ExpectErr(st.Float64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	When("complex type", func() {
		Specify("Complex64", func() {
			c := complex(float32(12), float32(13))
			r := 12
			ts := []*Table{
				New(int(r)),
				New(int8(r)),
				New(int16(r)),
				New(int32(r)),
				New(int64(r)),
				New(uint(r)),
				New(uint8(r)),
				New(uint16(r)),
				New(uint32(r)),
				New(uint64(r)),
				New(float32(r)),
			}
			for _, t := range ts {
				Expect(t.Complex64()).To(Equal(complex(float32(r), 0)))
			}
			tc := New(c)
			Expect(tc.Complex64()).To(Equal(c))

			// other type
			st := New("test")
			ExpectErr(st.Complex64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Complex128", func() {
			c := complex(float64(12), float64(13))
			r := 12
			ts := []*Table{
				New(int(r)),
				New(int8(r)),
				New(int16(r)),
				New(int32(r)),
				New(int64(r)),
				New(uint(r)),
				New(uint8(r)),
				New(uint16(r)),
				New(uint32(r)),
				New(uint64(r)),
				New(float32(r)),
				New(float64(r)),
			}
			for _, t := range ts {
				Expect(t.Complex128()).To(Equal(complex(float64(r), 0)))
			}
			tc := New(c)
			Expect(tc.Complex128()).To(Equal(c))

			// other type
			st := New("test")
			ExpectErr(st.Complex128()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	When("pack type", func() {
		Specify("Map", func() {
			m := map[int]int{
				1: 1,
				2: 2,
			}
			t := New(m)
			tm, err := t.Map()
			Expect(err).Should(BeNil())
			for tk, tv := range tm {
				Expect(m[tk.MustInt()]).Should(Equal(tv.MustInt()))
			}

			s := []int{1, 2}
			t = New(s)
			tm, err = t.Map()
			Expect(err).Should(BeNil())
			for tk, tv := range tm {
				Expect(s[tk.MustInt()]).Should(Equal(tv.MustInt()))
			}

			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			t = New(ss)
			tm, err = t.Map()
			Expect(err).Should(BeNil())
			for tk, tv := range tm {
				switch tk.String() {
				case "A":
					Expect(tv.Int()).Should(Equal(ss.A))
				case "B":
					Expect(tv.Int()).Should(Equal(ss.B))
				default:
					Expect(tk.String()).Should(BeEmpty())
				}
			}

			// other type
			st := New("test")
			ExpectErr(st.Map()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Slice", func() {
			s := []int{1, 2}
			t := New(s)
			ts, err := t.Slice()
			Expect(err).Should(BeNil())
			for idx, tv := range ts {
				Expect(tv.MustInt()).Should(Equal(s[idx]))
			}

			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			t = New(ss)
			ts, err = t.Slice()
			Expect(err).Should(BeNil())
			for idx, tv := range ts {
				switch idx {
				case 0:
					Expect(tv.MustInt()).Should(Equal(ss.A))
				case 1:
					Expect(tv.MustInt()).Should(Equal(ss.B))
				default:
					Expect(idx).Should(BeNumerically("<", 2))
				}
			}

			// other type
			st := New("test")
			ExpectErr(st.Slice()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
		Specify("Get", func() {
			m := map[int]int{
				1: 1,
				2: 2,
			}
			t := New(m)
			for k, v := range m {
				Expect(t.MustGet(k).Int()).Should(Equal(v))
			}
			Expect(t.MustGet(3)).Should(BeNil())

			s := []int{1, 2}
			t = New(s)
			for idx, elem := range s {
				Expect(t.MustGet(idx).Int()).Should(Equal(elem))
			}
			Expect(t.MustGet(3)).Should(BeNil())

			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			t = New(ss)
			Expect(t.MustGet("A").Int()).Should(Equal(ss.A))
			Expect(t.MustGet("B").Int()).Should(Equal(ss.B))
			Expect(t.MustGet("C")).Should(BeNil())

			// other type
			st := New("test")
			ExpectErr(st.Get("X")).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	When("each", func() {
		Specify("EachDo on map", func() {
			m := map[int]int{
				1: 1,
				2: 2,
			}
			t := New(m)
			t.EachDo(func(k, v *Table) error {
				Expect(m[k.MustInt()]).Should(Equal(v.MustInt()))
				return nil
			})
		})
	})
})
