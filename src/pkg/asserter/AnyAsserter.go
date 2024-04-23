package asserter

import (
	"assertG/src/pkg"
	"assertG/src/pkg/types"
	"reflect"
)

type AnyAsserter struct {
	t         types.T
	actual    any
	formatter *pkg.Formatter
}

func (anyA *AnyAsserter) IsEqualTo(expected any) *AnyAsserter {
	if !anyA.isEqual(expected) {
		anyA.mismatch(expected)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotEqualTo(expected any) *AnyAsserter {
	if anyA.isEqual(expected) {
		anyA.notExpected(expected)
	}
	return anyA
}

func (anyA *AnyAsserter) IsNil() *AnyAsserter {
	if anyA.actual != nil {
		anyA.mismatch("<nil>")
	}
	return anyA
}

func (anyA *AnyAsserter) IsNotNil() *AnyAsserter {
	if anyA.actual == nil {
		anyA.notExpected("<nil>")
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

func (anyA *AnyAsserter) notExpected(value any) {
	anyA.t.Fatalf(anyA.formatter.FormatMatch(value))
}

func (anyA *AnyAsserter) mismatch(expected any) {
	anyA.t.Fatalf(anyA.formatter.FormatMismatch(anyA.actual, expected))
}

// todo: convert formatter to explicit dependency

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual, formatter: pkg.NewFormatter()}
}
