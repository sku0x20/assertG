package pkg

import "assertG/src/pkg/types"

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

func NewStringAsserter(t types.T, actual string) *StringAsserter {
	return &StringAsserter{t: t, actual: actual, aa: NewAnyAsserter(t, actual)}
}
