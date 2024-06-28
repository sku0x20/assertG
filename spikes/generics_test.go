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

func Test_generics(t *testing.T) {
	makeStringVal()
	makeIntVal()
}
