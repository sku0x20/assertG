package runner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type TestFunc[T any] func(t *testing.T, extra T)
type SetupFunc[T any] func(t *testing.T) T
type TeardownFunc func(t *testing.T)

func NewTestsRunner[T any](t *testing.T) *TestsRunner[T] {
	return &TestsRunner[T]{
		t:        t,
		tests:    make([]TestFunc[T], 0, 10),
		setup:    func(t *testing.T) T { return *new(T) },
		teardown: func(t *testing.T) {},
	}
}

type TestsRunner[T any] struct {
	t        *testing.T
	tests    []TestFunc[T]
	setup    SetupFunc[T]
	teardown TeardownFunc
}

func (r *TestsRunner[T]) Add(f TestFunc[T]) {
	r.tests = append(r.tests, f)
}

func (r *TestsRunner[T]) Run() {
	for _, tf := range r.tests {
		r.t.Run(funcName(tf), func(t *testing.T) {
			extra := r.setup(t)
			tf(t, extra)
			r.teardown(t)
		})
	}
}

func (r *TestsRunner[T]) Setup(f SetupFunc[T]) {
	r.setup = f
}

func (r *TestsRunner[T]) Teardown(f TeardownFunc) {
	r.teardown = f
}

func funcName(f any) string {
	absoluteName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.Split(absoluteName, ".")[1]
}
