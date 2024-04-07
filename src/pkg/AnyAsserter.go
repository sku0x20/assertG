package pkg

import (
	"assertG/src/pkg/types"
	"fmt"
	"reflect"
)

type AnyAsserter struct {
	t      types.T
	actual any
}

func (anyA *AnyAsserter) IsEqualTo(expected any) *AnyAsserter {
	eqT, ok := anyA.actual.(types.Equals)
	isEqual := false
	if ok {
		isEqual = eqT.Equals(expected)
	} else {
		isEqual = reflect.DeepEqual(anyA.actual, expected)
	}
	if !isEqual {
		anyA.printError(expected)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotEqualTo(expected any) *AnyAsserter {
	eqT, ok := anyA.actual.(types.Equals)
	isEqual := false
	if ok {
		isEqual = eqT.Equals(expected)
	} else {
		isEqual = reflect.DeepEqual(anyA.actual, expected)
	}
	if isEqual {
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

func (anyA *AnyAsserter) printNotError(expected any) {
	anyA.printError(fmt.Sprintf("not %v", expected))
}

func (anyA *AnyAsserter) printError(expected any) {
	anyA.t.Fatalf("expected %v, got %v", expected, anyA.actual)
}

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual}
}
