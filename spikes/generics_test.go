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

func Test_generics3(t *testing.T) {
	sg := &SG[string]{}
	sg.Setup(func() string {
		return "test"
	})
	sg.Run(func(s string) {
		t.Logf("got %s", s)
	})
}

type SG[T any] struct {
	setup func() T
	run   func(T)
}

func (s *SG[T]) Setup(f func() T) {
	s.setup = f
}

func (s *SG[T]) Run(f func(T)) {
	r := s.setup()
	f(r)
}
