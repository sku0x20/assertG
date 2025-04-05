package asserter

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/assertG/src/main/message/verbs"
	"reflect"
)

func NewGeneric[T any](
	a assert_type.AssertType,
	equator equator.Equator[T],
	value T,
) *Generic[T] {
	return &Generic[T]{assertion: a, equator: equator, actual: value}
}

type Generic[T any] struct {
	assertion assert_type.AssertType
	equator   equator.Equator[T]
	actual    T
}

func (g *Generic[T]) IsEqualTo(val any) *Generic[T] {
	if !g.equator.AreEqual(g.actual, val) {
		g.assertion.FailWith(
			message.Expected().
				Value(g.actual).
				Verb(verbs.ToEqual).
				Value(val),
		)
	}
	return g
}

func (g *Generic[T]) IsNil() *Generic[T] {
	if !reflect.DeepEqual(g.actual, nil) {
		g.assertion.FailWith(
			message.Expected().
				Value(g.actual).
				Verb(verbs.ToBeNil),
		)
	}
	return g
}

func (g *Generic[T]) IsNotNil() *Generic[T] {
	if reflect.DeepEqual(g.actual, nil) {
		g.assertion.FailWith(
			message.Expected().
				Value(g.actual).
				Verb(verbs.NotToBeNil),
		)
	}
	return g
}
