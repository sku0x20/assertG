package assert_type

import (
	"github.com/sku0x20/assertG/src/main/message"
	"testing"
)

func NewHardAssert(t *testing.T) *HardAssert {
	return &HardAssert{t: t}
}

type HardAssert struct {
	t *testing.T
}

func (h *HardAssert) FailWith(msg *message.Message) {
	h.t.Fatalf("%s", msg.ToString())
}
