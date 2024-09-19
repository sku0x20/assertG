package assert

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
)

func (a *Assert) ThatAny(value any) *asserter.Any {
	return ThatAny(a.a, value)
}

func ThatAny(a assertion.Assertion, value any) *asserter.Any {
	return asserter.NewAny(a, value)
}
