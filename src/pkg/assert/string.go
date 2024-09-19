package assert

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
)

func (a *Assert) ThatString(value string) *asserter.String {
	return ThatString(a.a, value)
}

func ThatString(a assertion.Assertion, value string) *asserter.String {
	return asserter.NewString(a, value)
}
