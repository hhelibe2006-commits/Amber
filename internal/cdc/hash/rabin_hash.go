package hash

import "container/list"

type RabinHash struct {
	windows *list.List
	w       uint64
	p       uint64
	value   uint64
}

func NewRabinHash() *RabinHash {
	h := new(RabinHash)
	h.w = 64
	h.p = 257
	h.value = 0
	h.windows = list.New()
	return h
}

func (h *RabinHash) Next(value byte) {
	if h.w == uint64(h.windows.Len()) {
		h.value = (h.value-h.windows.Back().Value.(uint64)*pow(h.p, h.w-1))*h.p + uint64(value)
		h.windows.Remove(h.windows.Back())
		h.windows.PushFront(uint64(value))
	} else {
		h.value = h.value*h.p + uint64(value)
		h.windows.PushFront(uint64(value))
	}
}

func (h *RabinHash) Reset() {
	h.value = 0
	h.windows.Init()
}

func (h *RabinHash) Check(mask uint64) bool {
	return (h.value & mask) == 0
}

func pow(a, b uint64) uint64 {
	l := uint64(1)
	for b > 0 {
		b -= 1
		l *= a
	}
	return l
}
