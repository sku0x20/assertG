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
	case times == EXACTLY && order == ANY_ORDER:
		s.containsExactlyInAnyOrder(elems...)
	case times == ONCE && order == IN_ORDER:
		s.containsOnceInOrder(elems...)
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

func (s *Slice[T]) containsExactlyInAnyOrder(elems ...T) {
	if len(s.e) != len(elems) {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToContain).
				Value(elems).
				Verb(verbs.Exactly).
				Verb(verbs.InAnyOrder),
		)
		return
	}
	for _, e := range elems {
		if len(allIndexDeepEqual(s.e, e)) != 1 {
			s.a.FailWith(
				message.Expected().
					Value(s.e).
					Verb(verbs.ToContain).
					Value(elems).
					Verb(verbs.Exactly).
					Verb(verbs.InAnyOrder),
			)
			return
		}
	}
}

func (s *Slice[T]) containsOnceInOrder(elems ...T) {
	ei := make([]int, 0)
	for _, e := range elems {
		indexes := allIndexDeepEqual(s.e, e)
		if len(indexes) != 1 {
			s.a.FailWith(
				message.Expected().
					Value(s.e).
					Verb(verbs.ToContain).
					Value(elems).
					Verb(verbs.Once).
					Verb(verbs.InAnyOrder),
			)
			return
		}
		ei = append(ei, indexes[0])
	}
	for i := 1; i < len(ei); i++ {
		if ei[i] < ei[i-1] {
			s.a.FailWith(
				message.Expected().
					Value(s.e).
					Verb(verbs.ToContain).
					Value(elems).
					Verb(verbs.Once).
					Verb(verbs.InAnyOrder),
			)
		}
	}
}

func allIndexDeepEqual[T any](s []T, elem T) []int {
	return allIndex(s, func(t T) bool {
		return reflect.DeepEqual(t, elem)
	})
}

func allIndex[T any](s []T, f func(t T) bool) []int {
	in := make([]int, 0, len(s))
	for i := range s {
		if f(s[i]) {
			in = append(in, i)
		}
	}
	return in
}
