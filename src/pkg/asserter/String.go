package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	"strings"
)

func NewString(a assertion.Assertion, val string) *String {
	return &String{a, val}
}

type String struct {
	a assertion.Assertion
	e string
}

func (s *String) IsEqualTo(val string) {
	if s.e != val {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToEqual).
				Value(val),
		)
	}
}

func (s *String) Contains(val string) {
	if !strings.Contains(s.e, val) {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToContain).
				Value(val),
		)
	}
}

func (s *String) NotContains(val string) {
	if strings.Contains(s.e, val) {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToContain).
				Value(val),
		)
	}
}

func (s *String) HasLength(l int) {
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
}
