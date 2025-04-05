package assert

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/registry"
)

func ThatBool(val bool) *asserter.Bool {
	at := registry.GlobalRegistryGetAssertType()
	return ThatBoolWith(at, val)
}

func ThatBoolWith(a assert_type.AssertType, val bool) *asserter.Bool {
	return asserter.NewBool(a, val)
}
