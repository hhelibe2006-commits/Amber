package storage

type ChunkStore map[string][]byte

func NewChunkStore() ChunkStore {
	return make(ChunkStore)
}

type HashSet map[string]struct{}

func NewHashSet() HashSet {
	return make(HashSet)
}
