package asserter

import (
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_AnyAsserterNew(tm *testing.T) {
	r := runner.NewTestsRunner[any](tm)
	r.Setup(func(t *testing.T) any { return nil })
	r.Teardown(func(t *testing.T, e any) {})
	r.Add(PrintT)
	r.Run()
}

func PrintT(t *testing.T, e any) {
	t.Logf("ran")
}
