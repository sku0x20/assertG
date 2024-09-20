package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
)

func NewBool(a assertion.Assertion, val bool) *Bool {
	return &Bool{a: a, e: val}
}

type Bool struct {
	a assertion.Assertion
	e bool
}

func (b *Bool) IsTrue() *Bool {
	if !b.e {
		b.a.FailWith(
			message.Expected().
				Verb(verbs.ToBeTrue),
		)
	}
	return b
}

func (b *Bool) IsFalse() *Bool {
	if b.e {
		b.a.FailWith(
			message.Expected().
				Verb(verbs.ToBeFalse),
		)
	}
	return b
}
