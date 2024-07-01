package runner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type (
	TestFunc     = func(t *testing.T)
	SetupFunc    = func(t *testing.T) any
	TeardownFunc = func(t *testing.T)
)

func NewTestsRunner(t *testing.T) *TestsRunner {
	return &TestsRunner{
		t:     t,
		tests: make([]TestFunc, 0, 10),
		setup: func(t *testing.T) any { return nil },
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
		r.t.Run(funcName(tf), tf)
	}
}

func (r *TestsRunner) Setup(f SetupFunc) {
	r.setup = f
}

func (r *TestsRunner) TearDown(f TeardownFunc) {
	r.teardown = f
}

func funcName(f any) string {
	absoluteName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.Split(absoluteName, ".")[1]
}
