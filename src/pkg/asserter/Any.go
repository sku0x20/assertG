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

func (a *Any) IsEqualTo(val any) *Any {
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
	return a
}

func (a *Any) IsNil() *Any {
	if a.e != nil {
		a.a.FailWith(
			message.Expected().
				Value(a.e).
				Verb(verbs.ToBeNil),
		)
	}
	return a
}

func (a *Any) IsNotNil() *Any {
	if a.e == nil {
		a.a.FailWith(
			message.Expected().
				Value(a.e).
				Verb(verbs.NotToBeNil),
		)
	}
	return a
}
