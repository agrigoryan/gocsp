package csp

import (
	"bytes"
	"github.com/bits-and-blooms/bitset"
	"strconv"
)

type Domain struct {
	values ValueSet
	bitmap *bitset.BitSet
}

func (d *Domain) Size() int {
	return int(d.bitmap.Count())
}

func (d *Domain) Value(idx int) Value {
	return d.values[idx]
}

func (d *Domain) Range(fn func(int, Value) bool) {
	for i := 0; i < len(d.values); i++ {
		if d.bitmap.Test(uint(i)) {
			if fn(i, d.values[i]) {
				return
			}
		}
	}
}

func (d *Domain) Filter(fn func(Value) bool) {
	for i := 0; i < len(d.values); i++ {
		if d.bitmap.Test(uint(i)) {
			if !fn(d.values[i]) {
				d.bitmap.Clear(uint(i))
			}
		}
	}
}

func (d *Domain) Set(idx int) {
	d.bitmap.Set(uint(idx))
}

func (d *Domain) Unset(idx int) {
	d.bitmap.Clear(uint(idx))
}

func (d *Domain) UnsetAllBut(idx int) {
	d.bitmap.ClearAll()
	d.bitmap.Set(uint(idx))
}

func (d *Domain) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('{')
	d.Range(func(i int, value Value) bool {
		buf.WriteString(strconv.Itoa(int(value)))
		return false
	})
	buf.WriteByte('}')
	return buf.String()
}

func (d *Domain) Clone() Domain {
	return Domain{
		values: d.values,
		bitmap: d.bitmap.Clone(),
	}
}

func NewDomain(values []Value) Domain {
	return Domain{
		values: values,
		bitmap: bitset.New(uint(len(values))).SetAll(),
	}
}
