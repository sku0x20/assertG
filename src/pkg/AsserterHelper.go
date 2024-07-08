package pkg

import (
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/types"
)

type AsserterHelper struct {
	t      types.T
	actual any
}

func (h *AsserterHelper) Actual() any {
	return h.actual
}

func (h *AsserterHelper) Formatter() *message.Message {
	return message.Expected().
		Value(h.actual)
}

func (h *AsserterHelper) Error(builder *message.Message) {
	h.t.Fatalf(builder.ToString())
}

func NewAssertHelper(t types.T, actual any) *AsserterHelper {
	return &AsserterHelper{t: t, actual: actual}
}
