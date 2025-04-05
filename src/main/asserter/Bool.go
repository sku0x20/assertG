package asserter

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/assertG/src/main/message/verbs"
)

func NewBool(a assert_type.AssertType, val bool) *Bool {
	return &Bool{a: a, e: val}
}

type Bool struct {
	a assert_type.AssertType
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
