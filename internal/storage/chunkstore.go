package storage

type ChunkStore struct {
	Hash string
	Buf  []byte
}

func NewChunkStore() *ChunkStore {
	return new(ChunkStore)
}

type HashSet map[string]struct{}

func NewHashSet() HashSet {
	return make(HashSet)
}
