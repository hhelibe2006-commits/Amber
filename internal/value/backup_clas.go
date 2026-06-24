package value

type BackupClas struct {
	Typ    string
	Output string
	Input  PathList
	Cdc    string
}

func NewClas() *BackupClas {
	clas := new(BackupClas)
	return clas
}
