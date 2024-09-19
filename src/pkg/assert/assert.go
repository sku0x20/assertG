package assert

import "github.com/sku0x20/assertG/src/pkg/assertion"

func NewAssert(a assertion.Assertion) *Assert {
	return &Assert{a: a}
}

type Assert struct {
	a assertion.Assertion
}
