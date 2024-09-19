package assertion

import (
	"github.com/sku0x20/assertG/src/pkg/message"
	"testing"
)

func NewSoft(t *testing.T) *Soft {
	return &Soft{t, false}
}

type Soft struct {
	t      *testing.T
	failed bool
}

func (s *Soft) FailWith(msg *message.Message) {
	s.failed = true
	s.t.Helper()
	s.t.Logf(msg.ToString())
}

func (s *Soft) Failed() bool {
	return s.failed
}
