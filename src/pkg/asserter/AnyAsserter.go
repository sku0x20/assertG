package asserter

import (
	"assertG/src/pkg"
	"assertG/src/pkg/message"
	"assertG/src/pkg/message/messages"
	"assertG/src/pkg/types"
	"reflect"
)

type AnyAsserter struct {
	h *pkg.AsserterHelper
}

func (anyA *AnyAsserter) IsEqualTo(expected any) *AnyAsserter {
	if !anyA.isEqual(expected) {
		anyA.error(
			anyA.formatter().
				Message(messages.ToEqual).
				Value(expected),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotEqualTo(expected any) *AnyAsserter {
	if anyA.isEqual(expected) {
		anyA.error(
			anyA.formatter().
				Message(messages.NotToEqual).
				Value(expected),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNil() *AnyAsserter {
	if anyA.actual() != nil {
		anyA.error(
			anyA.formatter().
				Message(messages.ToBeNil),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotNil() *AnyAsserter {
	if anyA.actual() == nil {
		anyA.error(
			anyA.formatter().
				Message(messages.NotToBeNil),
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

func (anyA *AnyAsserter) error(builder *message.Builder) {
	anyA.h.Error(builder)
}

func (anyA *AnyAsserter) formatter() *message.Builder {
	return anyA.h.Formatter()
}

func NewAnyAsserter(h *pkg.AsserterHelper) *AnyAsserter {
	return &AnyAsserter{h: h}
}
