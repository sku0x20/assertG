package asserter

import (
	"assertG/src/pkg/types"
	"reflect"
)

type AnyAsserter struct {
	t      types.T
	actual any
}

func (anyA *AnyAsserter) IsEqualTo(expected any) *AnyAsserter {
	if !anyA.isEqual(expected) {
		anyA.error("test")
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotEqualTo(expected any) *AnyAsserter {
	if anyA.isEqual(expected) {
		anyA.error("test")
	}
	return anyA
}

func (anyA *AnyAsserter) IsNil() *AnyAsserter {
	if anyA.actual != nil {
		anyA.error("test")
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotNil() *AnyAsserter {
	if anyA.actual == nil {
		anyA.error("test")
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

func (anyA *AnyAsserter) error(str string) {
	anyA.t.Fatalf(str)
}

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual}
}
