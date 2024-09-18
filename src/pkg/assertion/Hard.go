package assertion

import (
	"github.com/sku0x20/assertG/src/pkg/api"
	"github.com/sku0x20/assertG/src/pkg/message"
)

func Hard(t api.T) *HardAssertion {
	return &HardAssertion{t: t}
}

type HardAssertion struct {
	t api.T
}

func (h *HardAssertion) FailWith(msg *message.Message) {
	h.t.Fatalf(msg.ToString())
}
