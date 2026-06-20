package hash

import "testing"

func TestPop(t *testing.T) {
	rabinHash := NewRabinHash()
	for i := 0; i < 256; i++ {
		if powTable[i] != uint64(i)*rabinHash.pow {
			t.Errorf("pop[%v] = %v,实际应为 %v", i, powTable[i], uint64(i)*rabinHash.pow)
		}
	}
}
