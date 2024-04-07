package pkg

import (
	"assertG/src/pkg/types"
	"reflect"
)

type GenericAsserter[T any] struct {
	t      types.T
	actual T
}

func (ga *GenericAsserter[T]) IsEqualTo(expected T) *GenericAsserter[T] {
	isEqual := reflect.DeepEqual(ga.actual, expected)
	if !isEqual {
		ga.t.Fatalf("expected %v, got %v", expected, ga.actual)
	}
	return ga
}

func (ga *GenericAsserter[T]) IsNil() *GenericAsserter[T] {
	if !reflect.ValueOf(ga.actual).IsNil() {
		ga.t.Fatalf("expected <nil>, got %v", ga.actual)
	}
	return ga
}

func NewGenericAsserter[T any](t types.T, actual T) *GenericAsserter[T] {
	return &GenericAsserter[T]{t: t, actual: actual}
}
