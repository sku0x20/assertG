package assertion

import (
	"github.com/sku0x20/assertG/src/pkg/message"
)

func NewFake() *Fake {
	return &Fake{}
}

type Fake struct {
}

func (f *Fake) FailWith(msg *message.Message) {
	//TODO implement me
	panic("implement me")
}
