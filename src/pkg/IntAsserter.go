package pkg

import (
	"assertG/src/pkg/types"
)

type IntAsserter struct {
	t      types.T
	actual int
}

func (a *IntAsserter) IsEqualTo(expected int) {
	if a.actual != expected {
		a.t.Fatalf("expected %d, but got %d", expected, a.actual)
	}
}

func NewAsserter(t types.T, actual int) *IntAsserter {
	if t == nil {
		panic("t cannot be nil")
	}
	return &IntAsserter{t, actual}
}
