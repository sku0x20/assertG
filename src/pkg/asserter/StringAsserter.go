package asserter

import (
	"assertG/src/pkg"
	"assertG/src/pkg/types"
	"strings"
)

type StringAsserter struct {
	t      types.T
	actual string
	aa     *AnyAsserter
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
	if !strings.Contains(sa.actual, expected) {
		sa.t.Fatalf("expected contains '%s', got '%s'", expected, sa.actual)
	}
	return sa
}

func (sa *StringAsserter) NotContains(expected string) *StringAsserter {
	if strings.Contains(sa.actual, expected) {
		sa.t.Fatalf("expected not contains '%s', got '%s'", expected, sa.actual)
	}
	return sa
}

func (sa *StringAsserter) HasLength(expected int) *StringAsserter {
	if len(sa.actual) != expected {
		sa.t.Fatalf("expected length '%d', got '%d'", expected, len(sa.actual))
	}
	return sa
}

func NewStringAsserter(t types.T, formatter *pkg.Formatter, actual string) *StringAsserter {
	return &StringAsserter{t: t, actual: actual, aa: NewAnyAsserter(t, formatter, actual)}
}
