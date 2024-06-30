package csp

import (
	"bytes"
	"github.com/bits-and-blooms/bitset"
	"strconv"
)

type Value int
type ValueSet []Value

type DomainWithBitmap struct {
	values ValueSet
	bitmap *bitset.BitSet
}

func (d *DomainWithBitmap) Size() int {
	return int(d.bitmap.Count())
}

func (d *DomainWithBitmap) Value(idx int) Value {
	return d.values[idx]
}

func (d *DomainWithBitmap) Range(fn func(int) bool) {
	for i := 0; i < len(d.values); i++ {
		if d.bitmap.Test(uint(i)) && fn(i) {
			return
		}
	}
}

func (d *DomainWithBitmap) Filter(fn func(int) bool) {
	for i := 0; i < len(d.values); i++ {
		if d.bitmap.Test(uint(i)) && !fn(i) {
			d.bitmap.Clear(uint(i))
		}
	}
}

func (d *DomainWithBitmap) Set(idx int) {
	d.bitmap.Set(uint(idx))
}

func (d *DomainWithBitmap) Unset(idx int) {
	d.bitmap.Clear(uint(idx))
}

func (d *DomainWithBitmap) Contains(idx int) bool {
	return d.bitmap.Test(uint(idx))
}

func (d *DomainWithBitmap) UnsetAllBut(idx int) {
	d.bitmap.ClearAll()
	d.bitmap.Set(uint(idx))
}

func (d *DomainWithBitmap) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('{')
	d.Range(func(i int) bool {
		buf.WriteString(strconv.Itoa(int(d.values[i])))
		return false
	})
	buf.WriteByte('}')
	return buf.String()
}

func (d *DomainWithBitmap) Clone() DomainWithBitmap {
	return DomainWithBitmap{
		values: d.values,
		bitmap: d.bitmap.Clone(),
	}
}

func NewDomainWithBitmap(values ValueSet) DomainWithBitmap {
	return DomainWithBitmap{
		values: values,
		bitmap: bitset.New(uint(len(values))).SetAll(),
	}
}

type DomainWithRemainingIndices struct {
	values    []Value
	remaining []int
}

func (d *DomainWithRemainingIndices) Size() int {
	return len(d.remaining)
}

func (d *DomainWithRemainingIndices) Value(idx int) Value {
	return d.values[idx]
}

func (d *DomainWithRemainingIndices) Range(fn func(int) bool) {
	for i := 0; i < len(d.remaining); i++ {
		if fn(d.remaining[i]) {
			return
		}
	}
}

func (d *DomainWithRemainingIndices) Filter(fn func(int) bool) {
	for i := 0; i < len(d.remaining); i++ {
		if !fn(d.remaining[i]) {
			d.Remove(d.remaining[i])
			i--
		}
	}
}

func (d *DomainWithRemainingIndices) Remove(idx int) {
	rIdx := -1
	for i := 0; i < len(d.remaining); i++ {
		if d.remaining[i] == idx {
			rIdx = i
			break
		}
	}
	if rIdx == -1 {
		return
	}
	d.remaining[len(d.remaining)-1], d.remaining[rIdx] = d.remaining[rIdx], d.remaining[len(d.remaining)-1]
	d.remaining = d.remaining[:len(d.remaining)-1]
}

func (d *DomainWithRemainingIndices) Contains(idx int) bool {
	for i := 0; i < len(d.remaining); i++ {
		if d.remaining[i] == idx {
			return true
		}
	}
	return false
}

func (d *DomainWithRemainingIndices) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('{')
	d.Range(func(i int) bool {
		buf.WriteString(strconv.Itoa(int(d.values[i])))
		return false
	})
	buf.WriteByte('}')
	return buf.String()
}

func (d *DomainWithRemainingIndices) Clone() DomainWithRemainingIndices {
	return DomainWithRemainingIndices{
		values:    d.values,
		remaining: d.remaining,
	}
}

func NewDomainWithRemainingIndices(values ValueSet) DomainWithRemainingIndices {
	remaining := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		remaining[i] = i
	}
	return DomainWithRemainingIndices{
		values:    values,
		remaining: remaining,
	}
}

type Domain = DomainWithRemainingIndices

func NewDomain(values ValueSet) Domain {
	return NewDomainWithRemainingIndices(values)
}
