package test

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_WithoutRunner(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	assert.ThatAny(10).IsNotNil().IsEqualTo(10)
}

func Test_WithRunner(tm *testing.T) {
	r := runner.NewTestsRunner[any](tm)
	r.Setup(func(t *testing.T) any {
		defer pkg.EnableAsserts(tm)()
		return nil
	})
	r.Teardown(func(t *testing.T, extra any) {
		t.Logf("tear-down")
	})
	r.Add(func(t *testing.T, extra any) {
		assert.ThatAny(10).IsNotNil().IsEqualTo(10)
	})
	r.Run()
}
