package asserter

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/assertG/src/main/message/verbs"
)

func NewSlice[T any](
	a assert_type.AssertType,
	elemEquator equator.Equator[T],
	val []T,
) *Slice[T] {
	return &Slice[T]{a, elemEquator, val}
}

type Slice[T any] struct {
	assertion   assert_type.AssertType
	elemEquator equator.Equator[T]
	actual      []T
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
	for i := 0; i < len(expected); i++ {
		if !s.elemEquator.AreEqual(s.actual[i], expected[i]) {
			s.assertion.FailWith(
				message.Expected().
					Value(s.actual).
					Verb(verbs.ToEqual).
					Value(expected),
			)
		}
	}
	return s
}

func (s *Slice[T]) IsEqualToIgnoringOrder(expected []T) *Slice[T] {
	s.HasLength(len(expected))
outer:
	for _, elem := range expected {
		for _, actual := range s.actual {
			if s.elemEquator.AreEqual(actual, elem) {
				continue outer
			}
		}
		s.assertion.FailWith(
			message.Expected().
				Value(s.actual).
				Verb(verbs.ToEqual).
				Value(expected),
		)
	}
	return s
}
