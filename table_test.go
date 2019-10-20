package table

import (
	"fmt"
	"math/bits"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ExpectErr(rets ...interface{}) Assertion {
	return Expect(rets[1])
}

var _ = Describe("Gets", func() {
	Context("with Bool()", func() {
		Specify("from bool kind", func() {
			b := true
			t := New(b)
			Expect(t.Bool()).To(Equal(b))
		})
		Specify("from bool ptr kind", func() {
			b := true
			t := New(&b)
			Expect(t.Bool()).To(Equal(b))
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Bool()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Bytes()", func() {
		Specify("from []byte type", func() {
			x := []byte("abcd")
			t := New(x)
			Expect(t.Bytes()).To(Equal(x))
		})
		Specify("from []byte ptr type", func() {
			x := []byte("abcd")
			t := New(&x)
			Expect(t.Bytes()).To(Equal(x))
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Bytes()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))

			t = New([]int{1, 2, 3})
			ExpectErr(t.Bytes()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Uint()", func() {
		Specify("from uint type", func() {
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
		})
		Specify("from ptr kind", func() {
			x := uint(1)
			t := New(&x)
			Expect(t.Uint()).Should(Equal(x))
		})
		Specify("from other type", func() {
			t := New("test")
			ExpectErr(t.Uint()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Uint8()", func() {
		Specify("from uint8 kind", func() {
			x := uint8(12)
			ts := []*Table{
				New(uint8(x)),
			}
			for _, t := range ts {
				Expect(t.Uint8()).To(Equal(x))
			}
		})
		Specify("from ptr kind", func() {
			x := uint8(1)
			t := New(&x)
			Expect(t.Uint8()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Uint8()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Uint16()", func() {
		Specify("from uint16 kind", func() {
			x := uint16(12)
			ts := []*Table{
				New(uint8(x)),
				New(uint16(x)),
			}
			for _, t := range ts {
				Expect(t.Uint16()).To(Equal(x))
			}
		})
		Specify("from ptr kind", func() {
			x := uint16(12)
			t := New(&x)
			Expect(t.Uint16()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Uint16()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Uint32()", func() {
		Specify("from uint32 kind", func() {
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
		})
		Specify("from ptr kind", func() {
			x := uint32(12)
			t := New(&x)
			Expect(t.Uint32()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Uint32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Uint64()", func() {
		Specify("from uint* kind", func() {
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
		})
		Specify("from ptr kind", func() {
			x := uint64(12)
			t := New(&x)
			Expect(t.Uint64()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Uint64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Int()", func() {
		Specify("from int, int8, int16, int32, int64, uint8, uint16, uint32 kind", func() {
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
		})
		Specify("from ptr kind", func() {
			x := int(12)
			t := New(&x)
			Expect(t.Int()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Int()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Int8()", func() {
		Specify("from int8 kind", func() {
			t8 := New(int8(12))
			Expect(t8.Int8()).To(Equal(int8(12)))
		})
		Specify("from ptr kind", func() {
			x := int8(12)
			t := New(&x)
			Expect(t.Int8()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			s := New("test")
			ExpectErr(s.Int8()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Int16()", func() {
		Specify("from int8, int16, uint8 kind", func() {
			x := int16(12)
			ts := []*Table{
				New(int8(x)),
				New(int16(x)),
				New(uint8(x)),
			}
			for _, t := range ts {
				Expect(t.Int16()).To(Equal(x))
			}
		})
		Specify("from ptr kind", func() {
			x := int16(12)
			t := New(&x)
			Expect(t.Int16()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Int16()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Int32()", func() {
		Specify("from int, int{8,16,32}, uint{8, 16}", func() {
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
		})
		Specify("from ptr kind", func() {
			x := int32(12)
			t := New(&x)
			Expect(t.Int32()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Int32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Int64()", func() {
		Specify("from int, int{8,16,32,64}, uint, uint{8, 16, 32}", func() {
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
		})
		Specify("from ptr kind", func() {
			x := int64(12)
			t := New(&x)
			Expect(t.Int64()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Int64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Float32()", func() {
		Specify("from int*, uint*, float32 kind", func() {
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
		})
		Specify("from ptr kind", func() {
			x := float32(12)
			t := New(&x)
			Expect(t.Float32()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Float32()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Float64()", func() {
		Specify("from int*, uint*, float* kind", func() {
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
		})
		Specify("from ptr kind", func() {
			x := float64(12)
			t := New(&x)
			Expect(t.Float64()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Float64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Complex64()", func() {
		Specify("from int*, uint*, float*, complex64 kind", func() {
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
		})
		Specify("from ptr kind", func() {
			x := complex(float32(12), float32(13))
			t := New(&x)
			Expect(t.Complex64()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Complex64()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Complex128()", func() {
		Specify("from int*, uint*, float*, complex* kind", func() {
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

			c32 := complex(float32(12), float32(13))
			tc32 := New(c32)
			Expect(tc32.Complex128()).To(Equal(c))
		})
		Specify("from ptr kind", func() {
			x := complex(float64(12), float64(13))
			t := New(&x)
			Expect(t.Complex128()).Should(Equal(x))
		})
		Specify("from other kind", func() {
			st := New("test")
			ExpectErr(st.Complex128()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Get()", func() {
		Specify("from map kind", func() {
			m := map[int]int{
				1: 1,
				2: 2,
			}
			t := New(m)
			for k, v := range m {
				Expect(t.MustGet(k).Int()).Should(Equal(v))
			}
			Expect(t.MustGet(3)).Should(BeNil())
		})
		Specify("from slice kind", func() {
			s := []int{1, 2}
			t := New(s)
			for idx, elem := range s {
				Expect(t.MustGet(idx).Int()).Should(Equal(elem))
			}
			Expect(t.MustGet(3)).Should(BeNil())
		})
		Specify("from array kind", func() {
			s := [3]int{1, 2}
			t := New(s)
			for idx, elem := range s {
				Expect(t.MustGet(idx).Int()).Should(Equal(elem))
			}
			Expect(t.MustGet(4)).Should(BeNil())
		})
		Specify("from struct kind", func() {
			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			t := New(ss)
			Expect(t.MustGet("A").Int()).Should(Equal(ss.A))
			Expect(t.MustGet("B").Int()).Should(Equal(ss.B))
			Expect(t.MustGet("C")).Should(BeNil())
		})
		Specify("from ptr kind", func() {
			s := []int{1}
			t := New(&s)
			Expect(t.MustGet(0).Int()).Should(Equal(1))
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Get("x")).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Map()", func() {
		Specify("from map kind", func() {
			m := map[int]int{
				1: 1,
				2: 2,
			}
			tm := New(m).MustMap()
			Expect(len(tm)).Should(Equal(len(m)))
			for tk, tv := range tm {
				Expect(tv.Int()).Should(Equal(m[tk.MustInt()]))
			}
		})
		Specify("from slice kind", func() {
			s := []int{1, 2}
			tm := New(s).MustMap()
			Expect(len(tm)).Should(Equal(len(s)))
			for idx, elem := range tm {
				Expect(elem.Int()).Should(Equal(s[idx.MustInt()]))
			}
		})
		Specify("from array kind", func() {
			s := [3]int{1, 2}
			tm := New(s).MustMap()
			Expect(len(tm)).Should(Equal(len(s)))
			for idx, elem := range tm {
				Expect(elem.Int()).Should(Equal(s[idx.MustInt()]))
			}
		})
		Specify("from struct kind", func() {
			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			tm := New(ss).MustMap()
			Expect(len(tm)).Should(Equal(2))
			for tk, tv := range tm {
				key, err := tk.String()
				Expect(err).Should(BeNil())
				switch key {
				case "A":
					Expect(tv.Int()).Should(Equal(ss.A))
				case "B":
					Expect(tv.Int()).Should(Equal(ss.B))
				default: // must error
					Expect(tv.Int()).Should(BeNil())
				}
			}
		})
		Specify("from ptr kind", func() {
			s := []int{1, 2}
			tm := New(&s).MustMap()
			Expect(len(tm)).Should(Equal(len(s)))
			for idx, elem := range tm {
				Expect(elem.Int()).Should(Equal(s[idx.MustInt()]))
			}
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Map()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with Slice()", func() {
		Specify("from slice kind", func() {
			s := []int{1, 2}
			ts := New(s).MustSlice()
			Expect(len(ts)).Should(Equal(len(s)))
			for idx, elem := range ts {
				Expect(elem.Int()).Should(Equal(s[idx]))
			}
		})
		Specify("from array kind", func() {
			s := [3]int{1, 2}
			ts := New(s).MustSlice()
			Expect(len(ts)).Should(Equal(len(s)))
			for idx, elem := range ts {
				Expect(elem.Int()).Should(Equal(s[idx]))
			}
		})
		Specify("from struct kind", func() {
			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			ts := New(ss).MustSlice()
			Expect(len(ts)).Should(Equal(2))
			Expect(ts[0].Int()).Should(Equal(ss.A))
			Expect(ts[1].Int()).Should(Equal(ss.B))
		})
		Specify("from ptr kind", func() {
			s := []int{1, 2}
			ts := New(&s).MustSlice()
			Expect(len(ts)).Should(Equal(len(s)))
			for idx, elem := range ts {
				Expect(elem.Int()).Should(Equal(s[idx]))
			}
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.Slice()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with AList()", func() {
		Specify("from map kind", func() {
			m := map[int]int{
				1: 1,
				2: 2,
			}
			tl := New(m).MustAList()
			Expect(len(tl)).Should(Equal(len(m)))
			for _, kv := range tl {
				Expect(kv[1].Int() /* value */).Should(Equal(m[kv[0].MustInt() /* key */]))
			}
		})
		Specify("from slice kind", func() {
			s := []int{1, 2}
			tl := New(s).MustAList()
			Expect(len(tl)).Should(Equal(len(s)))
			for _, kv := range tl {
				Expect(kv[1].Int() /* value */).Should(Equal(s[kv[0].MustInt() /* idx */]))
			}
		})
		Specify("from array kind", func() {
			s := [3]int{1, 2}
			tl := New(s).MustAList()
			Expect(len(tl)).Should(Equal(len(s)))
			for _, kv := range tl {
				Expect(kv[1].Int() /* value */).Should(Equal(s[kv[0].MustInt() /* idx */]))
			}
		})
		Specify("from struct kind", func() {
			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			tl := New(ss).MustAList()
			Expect(len(tl)).Should(Equal(2))

			Expect(tl[0][0].String() /* field name */).Should(Equal("A"))
			Expect(tl[1][0].String() /* field name */).Should(Equal("B"))

			Expect(tl[0][1].Int() /* value */).Should(Equal(ss.A))
			Expect(tl[1][1].Int() /* value */).Should(Equal(ss.B))
		})
		Specify("from ptr kind", func() {
			s := []int{1, 2}
			tl := New(&s).MustAList()
			Expect(len(tl)).Should(Equal(len(s)))
			for _, kv := range tl {
				Expect(kv[1].Int() /* value */).Should(Equal(s[kv[0].MustInt() /* idx */]))
			}
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.AList()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with PList()", func() {
		Specify("from map kind", func() {
			m := map[int]int{
				1: 1,
				2: 2,
			}
			tl := New(m).MustPList()
			Expect(len(tl)).Should(Equal(2 * len(m)))
			for i := 0; i < len(tl)/2; i += 2 {
				k := tl[i]
				v := tl[i+1]
				Expect(v.Int()).Should(Equal(m[k.MustInt()]))
			}
		})
		Specify("from slice kind", func() {
			s := []int{1, 2}
			tl := New(s).MustPList()
			Expect(len(tl)).Should(Equal(2 * len(s)))
			for i := 0; i < len(tl); i += 2 {
				k := tl[i]
				v := tl[i+1]
				Expect(v.Int()).Should(Equal(s[k.MustInt()]))
			}
		})
		Specify("from array kind", func() {
			s := [3]int{1, 2}
			tl := New(s).MustPList()
			Expect(len(tl)).Should(Equal(2 * len(s)))
			for i := 0; i < len(tl); i += 2 {
				k := tl[i]
				v := tl[i+1]
				Expect(v.Int()).Should(Equal(s[k.MustInt()]))
			}
		})
		Specify("from struct kind", func() {
			ss := struct {
				A, B int
			}{
				A: 1,
				B: 2,
			}
			tl := New(ss).MustPList()
			Expect(len(tl)).Should(Equal(2 * 2))
			Expect(tl[0].String() /* field name */).Should(Equal("A"))
			Expect(tl[1].Int() /* value */).Should(Equal(ss.A))
			Expect(tl[2].String() /* field name */).Should(Equal("B"))
			Expect(tl[3].Int() /* value */).Should(Equal(ss.B))
		})
		Specify("from ptr kind", func() {
			s := []int{1, 2}
			tl := New(&s).MustPList()
			Expect(len(tl)).Should(Equal(2 * len(s)))
			for i := 0; i < len(tl); i += 2 {
				k := tl[i]
				v := tl[i+1]
				Expect(v.Int()).Should(Equal(s[k.MustInt()]))
			}
		})
		Specify("from other kind", func() {
			t := New("test")
			ExpectErr(t.PList()).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
	Context("with String()", func() {
		Specify("bool kind", func() {
			b := true
			tb := New(b)
			Expect(tb.String()).To(Equal("true"))

			b = false
			tb = New(b)
			Expect(tb.String()).To(Equal("false"))
		})
		Specify("int* kind", func() {
			x := 123
			t := New(x)
			Expect(t.String()).Should(Equal("123"))
		})
		Specify("uint* kind", func() {
			x := uint(123)
			t := New(x)
			Expect(t.String()).Should(Equal("123"))
		})
		Specify("float* kind", func() {
			x := 1.2
			t := New(x)
			Expect(t.String()).Should(Equal("1.2"))

			x = 1.2e+34
			t = New(x)
			Expect(t.String()).Should(Equal("1.2e+34"))
		})
		Specify("complex* kind", func() {
			x := 1 + 2i
			t := New(x)
			Expect(t.String()).Should(Equal("(1+2i)"))

			x = -1.2e+34i + 1.2e+34
			t = New(x)
			Expect(t.String()).Should(Equal("(1.2e+34-1.2e+34i)"))
		})
	})
})

var _ = Describe("Sets", func() {
	Context("with Set()", func() {
		Specify("int kind", func() {
			m := 123
			t := New(&m)
			err := t.Set(2)
			Expect(err).Should(BeNil())
			Expect(t.Int()).Should(Equal(2))

			p := &m
			n := 456
			t = New(&p)
			err = t.Set(&n)
			Expect(err).Should(BeNil())
			Expect(t.Int()).Should(Equal(456))
		})
		Specify("value can't set", func() {
			t := New(1)
			Expect(t.Set(2)).To(BeAssignableToTypeOf((*ErrCannotSet)(nil)))
		})
		Specify("value kind unequal", func() {
			x := 1
			tx := New(&x)
			Expect(tx.Set(1.2)).To(BeAssignableToTypeOf((*ErrTypeUnequal)(nil)))
		})
	})

	Context("with Put()", func() {
		Specify("to map kind", func() {
			x := map[string]interface{}{
				"A": 1,
				"B": "b",
			}

			tx := New(x)
			Expect(tx.Put("A", 2)).Should(BeNil())
			Expect(tx.Put("C", 1.2)).Should(BeNil())

			Expect(tx.MustGet("A").Int()).Should(Equal(2))
			Expect(tx.MustGet("B").String()).Should(Equal("b"))
			Expect(tx.MustGet("C").Float64()).Should(Equal(1.2))
		})
		Specify("to slice kind", func() {
			x := []interface{}{1, "b"}

			tx := New(x)
			Expect(tx.Put(0, 2)).Should(BeNil())
			Expect(tx.Put(2, 1.2)).Should(BeNil())

			Expect(tx.MustGet(0).Int()).Should(Equal(2))
			Expect(tx.MustGet(1).String()).Should(Equal("b"))
			Expect(tx.MustGet(2).Float64()).Should(Equal(1.2))
		})
		Specify("to array kind", func() {
			x := [3]interface{}{1, "b"}

			tx := New(&x)
			Expect(tx.Put(0, 2)).Should(BeNil())
			Expect(tx.Put(2, 1.2)).Should(BeNil())

			Expect(tx.MustGet(0).Int()).Should(Equal(2))
			Expect(tx.MustGet(1).String()).Should(Equal("b"))
			Expect(tx.MustGet(2).Float64()).Should(Equal(1.2))
		})
		Specify("to struct kind", func() {
			x := struct {
				A int
				B string
				C string
			}{1, "b", ""}

			tx := New(&x)
			Expect(tx.Put("A", 2)).Should(BeNil())
			Expect(tx.Put("C", "c")).Should(BeNil())

			Expect(tx.MustGet("A").Int()).Should(Equal(2))
			Expect(tx.MustGet("B").String()).Should(Equal("b"))
			Expect(tx.MustGet("C").String()).Should(Equal("c"))
		})
		Specify("to other kind", func() {
			tx := New("a")
			Expect(tx.Put("nil", "nil")).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))

			x := 123
			tx = New(&x)
			Expect(tx.Put("nil", "nil")).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
		})
	})
})

var _ = Describe("Dos", func() {
	Context("with EachDo()", func() {
		Specify("in map", func() {
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

var _ = Describe("ConvTo", func() {
	Specify("not ptr kind", func() {
		x := true
		var y bool

		tx := New(x)
		err := tx.ConvTo(y)
		Expect(err).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
	})
	Specify("bool kind", func() {
		x := true
		var y bool

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y).Should(Equal(x))
	})
	Specify("int kind", func() {
		x := 123
		var y int

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y).Should(Equal(x))
	})
	Specify("uint kind", func() {
		x := uint(123)
		y := uint(0)

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y).Should(Equal(uint(x)))
	})
	Specify("float kind", func() {
		x := 123.4
		y := 0.0

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y).Should(Equal(x))
	})
	Specify("complex kind", func() {
		x := 1i + 2
		y := 0i + 0

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y).Should(Equal(x))
	})
	Specify("string kind", func() {
		x := "abc"
		y := ""

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y).Should(Equal(x))
	})
	Specify("slice kind", func() {
		x := []interface{}{
			1, "a", 0.1,
		}

		y := make([]interface{}, 0)

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(len(y)).Should(Equal(len(x)))
		for i, e := range y {
			Expect(e).Should(Equal(x[i]))
		}
	})
	Specify("array kind", func() {
		x := [3]int{1, 2, 3}
		var y [3]int

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		for i, v := range x {
			fmt.Fprintf(GinkgoWriter, "%d, %v", i, y[i])
			Expect(y[i]).Should(Equal(v))
		}
	})
	Specify("map kind", func() {
		x := map[string]interface{}{
			"A": 1,
			"B": "a",
			"C": []int{1, 2, 3},
		}

		y := make(map[string]interface{})

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		for k, v := range x {
			fmt.Fprintf(GinkgoWriter, "%s, %v", k, y[k])
			Expect(y[k]).Should(Equal(v))
		}
	})
	Specify("to struct kind", func() {
		x := map[string]interface{}{
			"a": 1,
			"A": 11,

			"B": "b",
			"b": "bb",

			"C": "c",
			"c": "cc",
		}

		var y struct {
			A int    `table:"a"`
			B string `table:"_"`
			C string
		}

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y.A).Should(Equal(1))
		Expect(y.B).Should(Equal(""))
		Expect(y.C).Should(Equal("c"))
		fmt.Fprint(GinkgoWriter, y.A)
	})

	Specify("time.Duration type", func() {
		x := "1s"
		var y time.Duration

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y).Should(Equal(1 * time.Second))
	})

	Specify("time.Time type", func() {
		now := time.Now()
		x := now.Format(TimeLayout)
		fmt.Fprintln(GinkgoWriter, x)
		var y time.Time

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		Expect(y.Format(TimeLayout)).Should(Equal(x))
	})

	Specify("nest struct kind", func() {
		type xx struct {
			X int
			Y int
		}
		var y struct {
			A int
			B time.Duration
			C *xx
			D *time.Time
		}
		x := map[string]interface{}{
			"A": 1,
			"B": "1s",
			"D": "Mon Jan 2 15:04:05 -0700 MST 2006",
			"C": map[string]int{
				"X": 10,
				"Y": 11,
			},
		}

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).Should(BeNil())
		// Expect(y.A).Should(Equal(1))
		// Expect(y.B).Should(Equal(""))
		// Expect(y.C).Should(Equal("c"))
		fmt.Fprint(GinkgoWriter, y, y.C)
	})
	Specify("to chan kind", func() {
		x := map[string]interface{}{
			"A": 1,
			"B": "a",
		}
		var y chan int

		tx := New(x)
		err := tx.ConvTo(&y)
		Expect(err).To(BeAssignableToTypeOf((*ErrUnsupportedKind)(nil)))
	})
})

var _ = Describe("Musts", func() {
	Specify("with MustInt8()", func() {
		x := int8(1)
		t := New(x)
		Expect(t.MustInt8()).To(Equal(x))

		t = New("test")
		Expect(func() { t.MustInt8() }).Should(Panic())
	})
	Specify("with MustInt16()", func() {
		x := int16(1)
		t := New(x)
		Expect(t.MustInt16()).To(Equal(x))

		t = New("test")
		Expect(func() { t.MustInt16() }).Should(Panic())
	})
	Specify("with MustInt32()", func() {
		x := int32(1)
		t := New(x)
		Expect(t.MustInt32()).To(Equal(x))

		t = New("test")
		Expect(func() { t.MustInt32() }).Should(Panic())
	})
	Specify("with MustInt64()", func() {
		x := int64(1)
		t := New(x)
		Expect(t.MustInt64()).To(Equal(x))

		t = New("test")
		Expect(func() { t.MustInt64() }).Should(Panic())
	})
	Specify("with MustInt()", func() {
		x := int(1)
		t := New(x)
		Expect(t.MustInt()).To(Equal(x))

		t = New("test")
		Expect(func() { t.MustInt() }).Should(Panic())
	})
})
