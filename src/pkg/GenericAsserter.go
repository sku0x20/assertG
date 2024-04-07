package pkg

import (
	"assertG/src/pkg/types"
	"reflect"
)

type GenericAsserter[T any] struct {
	t      types.T
	actual T
}

// not working

//func (ga *GenericAsserter[T]) IsNotNil() {
//	if ga.actual == nil {
//		ga.t.Fatalf("%v should not be nil", ga.actual)
//	}
//}

func (ga *GenericAsserter[T]) IsEqualTo(expected T) {
	reflect.DeepEqual(ga.actual, expected)
	//if ga.actual == expected {
	//	ga.t.Fatalf("%v should not be nil", ga.actual)
	//}
}

func NewGenericAsserter[T any](t types.T, actual T) *GenericAsserter[T] {
	return &GenericAsserter[T]{t: t, actual: actual}
}
