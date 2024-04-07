package pkg

import (
	"assertG/src/pkg/types"
	"reflect"
)

type AnyAsserter struct {
	t      types.T
	actual any
}

func (ga *AnyAsserter) IsEqualTo(expected any) *AnyAsserter {
	isEqual := reflect.DeepEqual(ga.actual, expected)
	if !isEqual {
		ga.t.Fatalf("expected %v, got %v", expected, ga.actual)
	}
	return ga
}

func (ga *AnyAsserter) IsNil() *AnyAsserter {
	return ga.IsEqualTo(nil)
}

func NewAnyAsserter(t types.T, actual any) *AnyAsserter {
	return &AnyAsserter{t: t, actual: actual}
}
