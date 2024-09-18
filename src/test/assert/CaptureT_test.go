package assert

import (
	"github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_CaptureT(tm *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](tm)
	r.Add(capturesT)
	r.Run()
}

func capturesT(t *testing.T, _ any) {
	c := assert.NewCaptureT(t)
	if c.GetT() != t {
		t.Fatalf("capturesT(): t mismatch")
	}
}
