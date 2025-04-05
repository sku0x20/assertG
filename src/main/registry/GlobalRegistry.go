package registry

import "github.com/sku0x20/assertG/src/main/assert_type"

var (
	assertType assert_type.AssertType = nil
)

func GlobalRegistryGetAssertType() assert_type.AssertType {
	return assertType
}

func GlobalRegistrySetAssertType(at assert_type.AssertType) {
	assertType = at
}
