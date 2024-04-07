package pkg

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
		anyA.printError(expected)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotEqualTo(expected any) *AnyAsserter {
	if anyA.isEqual(expected) {
		anyA.printNotError(expected)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNil() *AnyAsserter {
	if anyA.actual != nil {
		anyA.printError("<nil>")
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotNil() *AnyAsserter {
	if anyA.actual == nil {
		anyA.printNotError("<nil>")
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

func (anyA *AnyAsserter) printNotError(expected any) {
	anyA.t.Fatalf("not expected '%v', got '%v'", expected, anyA.actual)
}

func (anyA *AnyAsserter) printError(expected any) {
	anyA.t.Fatalf("expected '%v', got '%v'", expected, anyA.actual)
}

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual}
}
