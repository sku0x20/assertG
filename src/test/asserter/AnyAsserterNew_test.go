package asserter

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_AnyAsserterNew(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Setup(func(t2 *testing.T) any { return nil })
	r.Teardown(func(t2 *testing.T, e any) {})
	r.Add(PrintT)
	r.Run()
}

func PrintT(t *testing.T, e any) {
	t.Logf("ran")
}
