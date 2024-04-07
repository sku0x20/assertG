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
	isEqual := reflect.DeepEqual(anyA.actual, expected)
	if !isEqual {
		anyA.printError(expected)
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
		anyA.printError("not <nil>")
	}
	return anyA
}

func (anyA *AnyAsserter) printError(expected any) {
	anyA.t.Fatalf("expected %v, got %v", expected, anyA.actual)
}

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual}
}
