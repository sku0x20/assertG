package runner

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type TestFunc = func(t *testing.T)

func NewTestsRunner(t *testing.T) *TestsRunner {
	return &TestsRunner{t: t, tests: make([]TestFunc, 0, 10)}
}

type TestsRunner struct {
	t     *testing.T
	tests []TestFunc
}

func (r *TestsRunner) Register(f TestFunc) {
	r.tests = append(r.tests, f)
}

func (r *TestsRunner) Run() {
	for _, tf := range r.tests {
		r.t.Run(funcName(tf), tf)
	}
}

func funcName(f any) string {
	absoluteName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return strings.Split(absoluteName, ".")[1]
}
