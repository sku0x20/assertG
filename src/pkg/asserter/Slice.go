package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	"reflect"
)

func NewSlice[T any](a assertion.Assertion, val []T) *Slice[T] {
	return &Slice[T]{a, val}
}

type Times string
type Order string

const (
	EXACTLY  Times = "EXACTLY"
	ONCE           = "ONCE"
	MULTIPLE       = "MULTIPLE"
)

const (
	IN_ORDER  Order = "IN_ORDER"
	ANY_ORDER       = "ANY_ORDER"
)

type Slice[T any] struct {
	a assertion.Assertion
	e []T
}

func (s *Slice[T]) HasLength(l int) *Slice[T] {
	if len(s.e) != l {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToHaveLength).
				Value(l).
				Verb(verbs.ButWas).
				Value(len(s.e)),
		)
	}
	return s
}

func (s *Slice[T]) Contains(times Times, order Order, elems ...T) *Slice[T] {
	switch {
	case times == EXACTLY && order == IN_ORDER:
		s.containsExactlyInOrder(elems...)
	}
	return s
}

func (s *Slice[T]) containsExactlyInOrder(elems ...T) {
	if !reflect.DeepEqual(s.e, elems) {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToContain).
				Value(elems).
				Verb(verbs.Exactly).
				Verb(verbs.InOrder),
		)
	}
}
