package uniview

import (
	"os"
	"time"
)

type Token struct {
	event chan Packet
	err   chan error
	Err   error
}

func (tok *Token) Wait() bool {
	select {
	case <-tok.event:
		return true
	case err := <-tok.err:
		tok.Err = err
		return false
	}
}

func (tok *Token) WaitTimeout(timeout time.Duration) bool {

	select {
	case <-time.After(timeout):
		tok.Err = os.ErrDeadlineExceeded
		return false
	case <-tok.event:
		close(tok.event)
		return true
	case err := <-tok.err:
		tok.Err = err
		return false
	}
}
