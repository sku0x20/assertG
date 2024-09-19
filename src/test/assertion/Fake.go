package assertion

import (
	"github.com/sku0x20/assertG/src/pkg/message"
)

func NewFake() *Fake {
	return &Fake{}
}

type Fake struct {
	failed bool
}

func (f *Fake) FailWith(msg *message.Message) {
	f.failed = true
}

func (f *Fake) Failed() bool {
	return f.failed
}
