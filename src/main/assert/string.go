package assert

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/registry"
)

func ThatString(value string) *asserter.String {
	at := registry.GlobalRegistryGetAssertType()
	return ThatStringWith(at, value)
}

func ThatStringWith(a assert_type.AssertType, value string) *asserter.String {
	return asserter.NewString(a, value)
}
