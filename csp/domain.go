package csp

import (
	"bytes"
	"github.com/bits-and-blooms/bitset"
	"strconv"
)

type Value int
type ValueSet []Value

type domainWithBitmap struct {
	values ValueSet
	bitmap *bitset.BitSet
}

func (d *domainWithBitmap) Size() int {
	return int(d.bitmap.Count())
}

func (d *domainWithBitmap) Value(idx int) Value {
	return d.values[idx]
}

func (d *domainWithBitmap) Range(fn func(int) bool) {
	for i := 0; i < len(d.values); i++ {
		if d.bitmap.Test(uint(i)) && fn(i) {
			return
		}
	}
}

func (d *domainWithBitmap) Filter(fn func(int) bool) {
	for i := 0; i < len(d.values); i++ {
		if d.bitmap.Test(uint(i)) && !fn(i) {
			d.bitmap.Clear(uint(i))
		}
	}
}

func (d *domainWithBitmap) Remove(idx int) {
	d.bitmap.Clear(uint(idx))
}

func (d *domainWithBitmap) Contains(idx int) bool {
	return d.bitmap.Test(uint(idx))
}

func (d *domainWithBitmap) UnsetAllBut(idx int) {
	d.bitmap.ClearAll()
	d.bitmap.Set(uint(idx))
}

func (d *domainWithBitmap) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('{')
	d.Range(func(i int) bool {
		buf.WriteString(strconv.Itoa(int(d.values[i])))
		return false
	})
	buf.WriteByte('}')
	return buf.String()
}

func (d *domainWithBitmap) Clone() domainWithBitmap {
	return domainWithBitmap{
		values: d.values,
		bitmap: d.bitmap.Clone(),
	}
}

func newDomainWithBitmap(values ValueSet) domainWithBitmap {
	return domainWithBitmap{
		values: values,
		bitmap: bitset.New(uint(len(values))).SetAll(),
	}
}

type domainWithRemainingIndices struct {
	values    []Value
	remaining []int
}

func (d *domainWithRemainingIndices) Size() int {
	return len(d.remaining)
}

func (d *domainWithRemainingIndices) Value(idx int) Value {
	return d.values[idx]
}

func (d *domainWithRemainingIndices) Range(fn func(int) bool) {
	for i := 0; i < len(d.remaining); i++ {
		if fn(d.remaining[i]) {
			return
		}
	}
}

func (d *domainWithRemainingIndices) Filter(fn func(int) bool) {
	for i := 0; i < len(d.remaining); i++ {
		if !fn(d.remaining[i]) {
			d.Remove(d.remaining[i])
			i--
		}
	}
}

func (d *domainWithRemainingIndices) Remove(idx int) {
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

func (d *domainWithRemainingIndices) Contains(idx int) bool {
	for i := 0; i < len(d.remaining); i++ {
		if d.remaining[i] == idx {
			return true
		}
	}
	return false
}

func (d *domainWithRemainingIndices) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('{')
	d.Range(func(i int) bool {
		buf.WriteString(strconv.Itoa(int(d.values[i])))
		return false
	})
	buf.WriteByte('}')
	return buf.String()
}

func (d *domainWithRemainingIndices) Clone() domainWithRemainingIndices {
	return domainWithRemainingIndices{
		values:    d.values,
		remaining: d.remaining,
	}
}

func newDomainWithRemainingIndices(values ValueSet) domainWithRemainingIndices {
	remaining := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		remaining[i] = i
	}
	return domainWithRemainingIndices{
		values:    values,
		remaining: remaining,
	}
}

type Domain = domainWithBitmap

func NewDomain(values ValueSet) Domain {
	return newDomainWithBitmap(values)
}
