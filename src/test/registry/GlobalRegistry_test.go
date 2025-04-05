package registry

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/registry"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_GlobalRegistry(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(setAssertType)
	r.Run()
}

func setAssertType(t *testing.T, _ any) {
	sa := assert_type.NewSoftAssert(t)
	registry.GlobalRegistrySetAssertType(sa)
	if registry.GlobalRegistryGetAssertType() != sa {
		t.Fatalf("different AssertType")
	}
}
