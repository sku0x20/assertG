package pkg

import (
	"assertG/src/pkg/format"
	"assertG/src/pkg/types"
)

type AsserterHelper struct {
	t      types.T
	actual any
}

func (h *AsserterHelper) Actual() any {
	return h.actual
}

func (h *AsserterHelper) Formatter() *format.Builder {
	return format.Expected().
		Value(h.actual)
}

func (h *AsserterHelper) Error(builder *format.Builder) {
	h.t.Fatalf(builder.ToString())
}

func NewAssertHelper(t types.T, actual any) *AsserterHelper {
	return &AsserterHelper{t: t, actual: actual}
}
