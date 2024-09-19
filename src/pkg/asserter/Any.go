package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/api"
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
	casted, ok := a.e.(api.Equal)
	var equal bool
	if ok {
		equal = casted.Equal(val)
	} else {
		equal = reflect.DeepEqual(a.e, val)
	}
	if !equal {
		a.a.FailWith(
			message.Expected().
				Value(a.e).
				Verb(verbs.ToEqual).
				Value(val),
		)
	}
}

func (a *Any) IsNil() {
	if a.e != nil {
		a.a.FailWith(
			message.Expected().
				Value(a.e).
				Verb(verbs.ToBeNil),
		)
	}
}

func (a *Any) IsNotNil() {
	if a.e == nil {
		a.a.FailWith(
			message.Expected().
				Value(a.e).
				Verb(verbs.NotToBeNil),
		)
	}
}
