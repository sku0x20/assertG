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
		anyA.expected(expected)
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
		anyA.expected("<nil>")
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

func (anyA *AnyAsserter) expected(value any) {
	anyA.error(anyA.formatter.FormatMismatch(anyA.actual, value))
}

func (anyA *AnyAsserter) notExpected(value any) {
	anyA.error(anyA.formatter.FormatMatch(value))
}

func (anyA *AnyAsserter) error(str string) {
	anyA.t.Fatalf(str)
}

// todo: convert formatter to explicit dependency

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual, formatter: pkg.NewFormatter()}
}
