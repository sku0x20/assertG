package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"github.com/sku0x20/assertG/src/pkg/equator"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
)

func NewAny(a assert_type.AssertType, equator equator.Equator[any], value any) *Any {
	return &Any{assertion: a, equator: equator, actual: value}
}

type Any struct {
	assertion assert_type.AssertType
	equator   equator.Equator[any]
	actual    any
}

func (a *Any) IsEqualTo(val any) *Any {
	if !a.equator.AreEqual(a.actual, val) {
		a.assertion.FailWith(
			message.Expected().
				Value(a.actual).
				Verb(verbs.ToEqual).
				Value(val),
		)
	}
	return a
}

func (a *Any) IsNil() *Any {
	if a.actual != nil {
		a.assertion.FailWith(
			message.Expected().
				Value(a.actual).
				Verb(verbs.ToBeNil),
		)
	}
	return a
}

func (a *Any) IsNotNil() *Any {
	if a.actual == nil {
		a.assertion.FailWith(
			message.Expected().
				Value(a.actual).
				Verb(verbs.NotToBeNil),
		)
	}
	return a
}
