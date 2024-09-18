package assertion

import (
	"github.com/sku0x20/assertG/src/pkg/message"
	"testing"
)

func NewHard(t *testing.T) *Hard {
	return &Hard{t: t}
}

type Hard struct {
	t *testing.T
}

func (h *Hard) FailWith(msg *message.Message) {
	h.t.Fatalf("%s", msg.ToString())
}
