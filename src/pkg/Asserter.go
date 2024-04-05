package pkg

import (
	"assertG/src/pkg/assert"
)

type Asserter struct {
	t      assert.T
	actual int
}

func (a *Asserter) IsEqualTo(expected int) {
	if a.actual != expected {
		a.t.Fatalf("expected %d, but got %d", expected, a.actual)
	}
}

func NewAsserter(t assert.T, actual int) *Asserter {
	return &Asserter{t, actual}
}
