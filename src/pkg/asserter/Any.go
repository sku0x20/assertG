package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
)

func NewAny(a assertion.Assertion, value any) *Any {
	return &Any{a: a}
}

type Any struct {
	a assertion.Assertion
	e any
}

func (a *Any) IsEqualTo(val any) {
	a.a.FailWith(
		message.Expected().
			Value(a.e).
			Verb(verbs.ToEqual),
	)
}
