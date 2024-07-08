package asserter

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	"strings"
)

type StringAsserter struct {
	h  *pkg.AsserterHelper
	aa *AnyAsserter
}

func (sa *StringAsserter) IsEqualTo(expected string) *StringAsserter {
	sa.aa.IsEqualTo(expected)
	return sa
}

func (sa *StringAsserter) IsNotEqualTo(expected string) *StringAsserter {
	sa.aa.IsNotEqualTo(expected)
	return sa
}

func (sa *StringAsserter) Contains(expected string) *StringAsserter {
	if !strings.Contains(sa.actual(), expected) {
		sa.error(
			sa.formatter().
				Verb(verbs.ToContain).
				Value(expected),
		)
	}
	return sa
}

func (sa *StringAsserter) NotContains(expected string) *StringAsserter {
	if strings.Contains(sa.actual(), expected) {
		sa.error(
			sa.formatter().
				Verb(verbs.NotToContain).
				Value(expected),
		)
	}
	return sa
}

func (sa *StringAsserter) HasLength(expected int) *StringAsserter {
	if len(sa.actual()) != expected {
		sa.error(
			sa.formatter().
				Verb(verbs.ToHaveLength).
				Value(expected).
				Verb(verbs.ButWas).
				Value(len(sa.actual())),
		)
	}
	return sa
}

func (sa *StringAsserter) actual() string {
	return sa.h.Actual().(string)
}

func (sa *StringAsserter) error(builder *message.Message) {
	sa.h.Error(builder)
}

func (sa *StringAsserter) formatter() *message.Message {
	return sa.h.Formatter()
}

func NewStringAsserter(h *pkg.AsserterHelper) *StringAsserter {
	return &StringAsserter{h: h, aa: NewAnyAsserter(h)}
}
