package storage

type ChunkStore map[string][]byte

func NewChunkStore() ChunkStore {
	return make(ChunkStore)
}
