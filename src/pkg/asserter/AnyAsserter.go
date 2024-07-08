package asserter

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	"github.com/sku0x20/assertG/src/pkg/types"
	"reflect"
)

type AnyAsserter struct {
	h *pkg.AsserterHelper
}

func (anyA *AnyAsserter) IsEqualTo(expected any) *AnyAsserter {
	if !anyA.isEqual(expected) {
		anyA.error(
			anyA.formatter().
				Verb(verbs.ToEqual).
				Value(expected),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotEqualTo(expected any) *AnyAsserter {
	if anyA.isEqual(expected) {
		anyA.error(
			anyA.formatter().
				Verb(verbs.NotToEqual).
				Value(expected),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNil() *AnyAsserter {
	if anyA.actual() != nil {
		anyA.error(
			anyA.formatter().
				Verb(verbs.ToBeNil),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotNil() *AnyAsserter {
	if anyA.actual() == nil {
		anyA.error(
			anyA.formatter().
				Verb(verbs.NotToBeNil),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) isEqual(expected any) bool {
	eqT, ok := anyA.actual().(types.Equals)
	if ok {
		return eqT.Equals(expected)
	}
	return reflect.DeepEqual(anyA.actual(), expected)
}

func (anyA *AnyAsserter) actual() any {
	return anyA.h.Actual()
}

func (anyA *AnyAsserter) error(builder *message.Message) {
	anyA.h.Error(builder)
}

func (anyA *AnyAsserter) formatter() *message.Message {
	return anyA.h.Formatter()
}

func NewAnyAsserter(h *pkg.AsserterHelper) *AnyAsserter {
	return &AnyAsserter{h: h}
}
