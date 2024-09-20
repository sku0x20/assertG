package assert

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
)

func (a *Assert) ThatBool(val bool) *asserter.Bool {
	return ThatBool(a.a, val)
}

func ThatBool(a assertion.Assertion, val bool) *asserter.Bool {
	return asserter.NewBool(a, val)
}
