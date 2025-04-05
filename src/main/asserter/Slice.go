package asserter

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/assertG/src/main/message/verbs"
	"reflect"
)

func NewSlice[T any](a assert_type.AssertType, val []T) *Slice[T] {
	return &Slice[T]{a, val}
}

type Slice[T any] struct {
	assertion assert_type.AssertType
	actual    []T
}

func (s *Slice[T]) HasLength(l int) *Slice[T] {
	if len(s.actual) != l {
		s.assertion.FailWith(
			message.Expected().
				Value(s.actual).
				Verb(verbs.ToHaveLength).
				Value(l).
				Verb(verbs.ButWas).
				Value(len(s.actual)),
		)
	}
	return s
}

func (s *Slice[T]) IsEqualTo(expected []T) *Slice[T] {
	s.HasLength(len(expected))
	if !reflect.DeepEqual(s.actual, expected) {
		s.assertion.FailWith(
			message.Expected().
				Value(s.actual).
				Verb(verbs.ToEqual).
				Value(expected),
		)
	}
	return s
}
