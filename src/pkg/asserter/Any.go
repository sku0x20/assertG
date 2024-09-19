package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	"reflect"
)

func NewAny(a assertion.Assertion, value any) *Any {
	return &Any{a: a, e: value}
}

type Any struct {
	a assertion.Assertion
	e any
}

func (a *Any) IsEqualTo(val any) {
	if !reflect.DeepEqual(a.e, val) {
		a.a.FailWith(
			message.Expected().
				Value(a.e).
				Verb(verbs.ToEqual).
				Value(val),
		)
	}
}
