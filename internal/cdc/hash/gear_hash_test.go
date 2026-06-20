package hash

import (
	"math/rand/v2"
	"testing"
)

func TestGear(t *testing.T) {
	pcg := rand.NewPCG(0, 0)
	for i := 0; i < 256; i++ {
		value := pcg.Uint64()
		if value != gear[i] {
			t.Errorf("gear[%d] = %d, 应为 %d", i, gear[i], value)
		}
	}
}
