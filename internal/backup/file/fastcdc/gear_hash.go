package fastcdc

import "math/rand/v2"

type GearHash struct {
	value uint64
	Gear  [256]uint64
}

func NewGearHash() *GearHash {
	g := new(GearHash)
	r := rand.New(rand.NewPCG(0, 0))
	g.value = 0
	for i := 0; i < len(g.Gear); i++ {
		g.Gear[i] = r.Uint64()
		println(g.Gear[i])
	}
	return g
}
