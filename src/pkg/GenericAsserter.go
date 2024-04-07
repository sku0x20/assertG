package pkg

import (
	"assertG/src/pkg/types"
	"reflect"
)

type GenericAsserter struct {
	t      types.T
	actual any
}

func (ga *GenericAsserter) IsEqualTo(expected any) *GenericAsserter {
	isEqual := reflect.DeepEqual(ga.actual, expected)
	if !isEqual {
		ga.t.Fatalf("expected %v, got %v", expected, ga.actual)
	}
	return ga
}

func (ga *GenericAsserter) IsNil() *GenericAsserter {
	return ga.IsEqualTo(nil)
}

func NewGenericAsserter[T any](t types.T, actual T) *GenericAsserter {
	return &GenericAsserter{t: t, actual: actual}
}
