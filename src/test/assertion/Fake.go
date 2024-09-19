package assertion

import (
	"github.com/sku0x20/assertG/src/pkg/message"
	"testing"
)

func NewFake(t *testing.T) *Fake {
	return &Fake{t, false}
}

type Fake struct {
	t      *testing.T
	failed bool
}

func (f *Fake) FailWith(msg *message.Message) {
	f.failed = true
	f.t.Helper()
	f.t.Logf(msg.ToString())
}

func (f *Fake) Failed() bool {
	return f.failed
}
