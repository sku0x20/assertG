package assertions

import (
	"assertG/src/pkg/message"
	"assertG/src/pkg/types"
)

func NewHardAssertion(t types.T) *HardAssertion {
	return &HardAssertion{t: t}
}

type HardAssertion struct {
	t types.T
}

func (h *HardAssertion) FailWith(msg *message.Message) {
	h.t.Fatalf("something")
}
