package pkg

import (
	"assertG/src/pkg/message"
	"assertG/src/pkg/types"
)

type AsserterHelper struct {
	t      types.T
	actual any
}

func (h *AsserterHelper) Actual() any {
	return h.actual
}

func (h *AsserterHelper) Formatter() *message.Builder {
	return message.Expected().
		Value(h.actual)
}

func (h *AsserterHelper) Error(builder *message.Builder) {
	h.t.Fatalf(builder.ToString())
}

func NewAssertHelper(t types.T, actual any) *AsserterHelper {
	return &AsserterHelper{t: t, actual: actual}
}
