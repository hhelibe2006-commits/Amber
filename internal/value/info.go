package value

type Info struct {
	Path  string
	Hash  [32]byte
	Value []byte
}

func (i *Info) Set(path string, value *[]byte, hash [32]byte, ch chan Info) {
	i.Path = path
	i.Hash = hash
	i.Value = (*value)[:len(*value):len(*value)]
	ch <- *i
	*value = (*value)[:0]
}
