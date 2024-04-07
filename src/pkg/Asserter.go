package pkg

import (
	"assertG/src/pkg/types"
)

type Asserter struct {
	t      types.T
	actual int
}

func (a *Asserter) IsEqualTo(expected int) {
	if a.actual != expected {
		a.t.Fatalf("expected %d, but got %d", expected, a.actual)
	}
}

func NewAsserter(t types.T, actual int) *Asserter {
	if t == nil {
		panic("t cannot be nil")
	}
	return &Asserter{t, actual}
}
