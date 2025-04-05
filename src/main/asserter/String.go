package asserter

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/assertG/src/main/message/verbs"
	"strings"
)

func NewString(a assert_type.AssertType, val string) *String {
	return &String{a, val}
}

type String struct {
	a assert_type.AssertType
	e string
}

func (s *String) IsEqualTo(val string) *String {
	if s.e != val {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToEqual).
				Value(val),
		)
	}
	return s
}

func (s *String) Contains(val string) *String {
	if !strings.Contains(s.e, val) {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToContain).
				Value(val),
		)
	}
	return s
}

func (s *String) NotContains(val string) *String {
	if strings.Contains(s.e, val) {
		s.a.FailWith(
			message.Expected().
				Value(s.e).
				Verb(verbs.ToContain).
				Value(val),
		)
	}
	return s
}

func (s *String) HasLength(l int) *String {
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
