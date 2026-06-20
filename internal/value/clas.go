package value

type Clas struct {
	Typ    string
	Output string
	Input  PathList
	Cdc    string
}

func NewClas() *Clas {
	clas := new(Clas)
	return clas
}
