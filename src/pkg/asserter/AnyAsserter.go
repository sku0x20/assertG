package asserter

import (
	"assertG/src/pkg/format"
	"assertG/src/pkg/types"
	"reflect"
)

type AnyAsserter struct {
	t      types.T
	actual any
}

func (anyA *AnyAsserter) IsEqualTo(expected any) *AnyAsserter {
	if !anyA.isEqual(expected) {
		anyA.error(
			anyA.formatter().
				Message(format.ToEqual).
				Value(expected),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotEqualTo(expected any) *AnyAsserter {
	if anyA.isEqual(expected) {
		anyA.error(
			anyA.formatter().
				Message(format.NotToEqual).
				Value(expected),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNil() *AnyAsserter {
	if anyA.actual != nil {
		anyA.error(
			anyA.formatter().
				Message(format.ToBeNil),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotNil() *AnyAsserter {
	if anyA.actual == nil {
		anyA.error(
			anyA.formatter().
				Message(format.NotToBeNil),
		)
	}
	return anyA
}

func (anyA *AnyAsserter) isEqual(expected any) bool {
	eqT, ok := anyA.actual.(types.Equals)
	if ok {
		return eqT.Equals(expected)
	}
	return reflect.DeepEqual(anyA.actual, expected)
}

func (anyA *AnyAsserter) error(builder *format.Builder) {
	anyA.t.Fatalf(builder.ToString())
}

func (anyA *AnyAsserter) formatter() *format.Builder {
	return format.Expected().
		Value(anyA.actual)
}

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual}
}
