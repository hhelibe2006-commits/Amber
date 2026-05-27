package fastcdc

// BlockMeta 用于文件切片的哈希值、字节切片以及在文件的位置
type BlockMeta struct {
	Hash  string
	bytes []byte
	Index uint64
}

// FastCDC 用于将单个文件切分成多个块，并将每个块的哈希值存储在chunkStore中，同时将文件的元数据存储在Manifest中
func FastCDC(input string, ch chan map[string][]byte) error {
	avg := 8 * 1024
	minByte, maxByte := avg/4, avg*4
	l, r := (1<<13)-1, (1<<11)-1
	println(minByte, maxByte, l, r)
	return nil
}
