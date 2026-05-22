package fastcdc

import "math/rand"

type GearHash struct {
	table [256]uint64
	l     uint64
}

func NewGearHash(i int64) *GearHash {
	g := new(GearHash)
	rng := rand.New(rand.NewSource(i))
	for i := 0; i < 256; i++ {
		g.table[i] = rng.Uint64()
	}
	return g
}

func FastCDC(i int64) {
	gearHash := NewGearHash(i)
	println(gearHash)
}
