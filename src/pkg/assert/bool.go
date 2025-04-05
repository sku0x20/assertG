package assert

import (
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"github.com/sku0x20/assertG/src/pkg/asserter"
)

func (a *Assert) ThatBool(val bool) *asserter.Bool {
	return ThatBool(a.a, val)
}

func ThatBool(a assert_type.AssertType, val bool) *asserter.Bool {
	return asserter.NewBool(a, val)
}
