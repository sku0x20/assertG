package test

import (
	"github.com/sku0x20/assertG/src/main/assert"
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/registry"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_WithoutRunner(t *testing.T) {
	registry.GlobalRegistrySetAssertType(assert_type.NewHardAssert(t))
	assert.That(10).
		IsNotNil().
		IsEqualTo(10)
}

func initE(t *testing.T) any {
	registry.GlobalRegistrySetAssertType(assert_type.NewSoftAssert(t))
	return nil
}

func Test_WithRunner(t *testing.T) {
	r := runner.NewTestsRunner[any](t, initE)
	r.Add(makeAndValidateAssert)
	r.Run()
}

func makeAndValidateAssert(t *testing.T, _ any) {
	assert.That(10).
		IsNotNil().
		IsEqualTo(20)

	sa := registry.GlobalRegistryGetAssertType().(*assert_type.SoftAssert)
	if !sa.Failed() {
		t.Fatalf("should have failed")
	}
}
