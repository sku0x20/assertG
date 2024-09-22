package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
)

func NewSlice[T any](a assertion.Assertion, val []T) *Slice[T] {
	return &Slice[T]{a, val}
}

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
