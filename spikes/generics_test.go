package spikes

import (
	"fmt"
	"testing"
)

type test[T any] struct {
	val T
}

func (t *test[T]) getVal() T {
	return t.val
}

func (t *test[T]) setVal(val T) {
	t.val = val
}

func makeStringVal() {
	tString := &test[string]{val: "s1"}
	fmt.Printf("%s\n", tString.getVal())
	tString.setVal("s2")
	fmt.Printf("%s\n", tString.getVal())
}

func makeIntVal() {
	tString := &test[int]{val: 23}
	fmt.Printf("%d\n", tString.getVal())
	tString.setVal(43)
	fmt.Printf("%d\n", tString.getVal())
}

func Test_generics1(t *testing.T) {
	makeStringVal()
	makeIntVal()
}

func Test_generics2(t *testing.T) {
	a := make([]any, 10)
	a[0] = 1
	a[1] = "r"

	t1 := maker("tr")
	//t1[0] = 1 // invalid
	t1[1] = "r"
	taker(a)
	taker(t1)
}

func maker[T any](t T) []T {
	return make([]T, 90)
}

func taker[T any](in []T) {
	// take
}
