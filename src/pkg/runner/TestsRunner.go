package runner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type TestFunc func(t *testing.T, extra any)
type SetupFunc func(t *testing.T) any
type TeardownFunc func(t *testing.T)

func NewTestsRunner(t *testing.T) *TestsRunner {
	return &TestsRunner{
		t:        t,
		tests:    make([]TestFunc, 0, 10),
		setup:    func(t *testing.T) any { return nil },
		teardown: func(t *testing.T) {},
	}
}

type TestsRunner struct {
	t        *testing.T
	tests    []TestFunc
	setup    SetupFunc
	teardown TeardownFunc
}

func (r *TestsRunner) Add(f TestFunc) {
	r.tests = append(r.tests, f)
}

func (r *TestsRunner) Run() {
	for _, tf := range r.tests {
		r.t.Run(funcName(tf), func(t *testing.T) {
			extra := r.setup(t)
			tf(t, extra)
			r.teardown(t)
		})
	}
}

func (r *TestsRunner) Setup(f SetupFunc) {
	r.setup = f
}

func (r *TestsRunner) Teardown(f TeardownFunc) {
	r.teardown = f
}

func funcName(f any) string {
	absoluteName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.Split(absoluteName, ".")[1]
}
