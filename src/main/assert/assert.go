package assert

import "github.com/sku0x20/assertG/src/main/assert_type"

func NewAssert(a assert_type.AssertType) *Assert {
	return &Assert{a: a}
}

type Assert struct {
	a assert_type.AssertType
}
