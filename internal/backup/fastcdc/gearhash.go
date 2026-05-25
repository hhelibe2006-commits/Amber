package fastcdc

import "math/rand/v2"

type GearHash struct {
	table     [256]uint64
	hashValue uint64
}

func (g *GearHash) Next(b byte) {
	g.hashValue = (g.hashValue << 1) + g.table[b]
}

func NewGearHash() *GearHash {
	g := new(GearHash)
	rng := rand.New(rand.NewPCG(0, 0))
	for i := 0; i < 256; i++ {
		g.table[i] = rng.Uint64()
	}
	return g
}
