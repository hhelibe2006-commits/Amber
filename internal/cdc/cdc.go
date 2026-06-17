package cdc

type Cdc interface {
	Next() ([]byte, error)
}
