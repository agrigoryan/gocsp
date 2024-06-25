package csp

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/kelindar/bitmap"
	"testing"
)

type MyBitsetInterface interface {
	Set(bit uint)
	Clear(bit uint)
	Count() uint
}

type MyBitset struct {
	set uint64
}

func (bs MyBitset) Set(bit uint) {
	if bit < 10 {
		bs.set |= 1 << bit
	} else {
		bs.set |= 1 << (bit + 1)
	}
}

func (bs MyBitset) Clear(bit uint) {
	bs.set &^= 1 << bit
}

func (bs MyBitset) Count() uint {
	return 5
}

func BenchmarkBitmap(b *testing.B) {
	bm := make(bitmap.Bitmap, 1)
	for n := 0; n < b.N; n++ {
		bm.Set(uint32(n % 32))
		//bm.Count()
		bm.Remove(5)
	}
}

func BenchmarkBitset(b *testing.B) {
	bs := bitset.New(60)
	for n := 0; n < b.N; n++ {
		bs.Set(uint(n % 32))
		//bs.Count()
		bs.Clear(5)
	}
}

func BenchmarkMyBitset(b *testing.B) {
	var bs MyBitset
	bs = MyBitset{0}
	for n := 0; n < b.N; n++ {
		bs.Set(uint(n % 32))
		bs.Count()
		bs.Clear(5)
	}
}
