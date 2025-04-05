package equator

import "reflect"

type ReflectDeepEquator[T any] struct {
}

func NewReflectDeepEquator[T any]() *ReflectDeepEquator[T] {
	return &ReflectDeepEquator[T]{}
}

func (r *ReflectDeepEquator[T]) AreEqual(a T, b any) bool {
	return reflect.DeepEqual(a, b)
}
