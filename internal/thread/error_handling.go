package thread

import "sync"

type ErrorHanding struct {
	Wg      *sync.WaitGroup
	Sem     chan struct{}
	ErrChan chan error
}

func NewErrorHanding(le uint8) *ErrorHanding {
	er := new(ErrorHanding)
	er.Wg = new(sync.WaitGroup)
	er.Sem = make(chan struct{}, le)
	er.ErrChan = make(chan error)
	return er
}
