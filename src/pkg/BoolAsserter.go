package pkg

import "assertG/src/pkg/types"

type BoolAsserter struct {
	t      types.T
	actual bool
}

func (a *BoolAsserter) IsNotNil() *BoolAsserter {
	return a
}

func NewBoolAsserter(t types.T, actual bool) *BoolAsserter {
	return &BoolAsserter{t, actual}
}
