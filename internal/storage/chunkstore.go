package storage

type ChunkStore map[[32]byte][]byte

func NewChunkStore() ChunkStore {
	return make(ChunkStore)
}
