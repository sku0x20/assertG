package asserter

import "github.com/sku0x20/assertG/src/pkg/assertion"

func NewAny(a assertion.Assertion) *Any {
	return &Any{a: a}
}

type Any struct {
	a assertion.Assertion
}
