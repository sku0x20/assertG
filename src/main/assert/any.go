package assert

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
)

func (a *Assert) ThatAny(value any) *asserter.Any {
	return ThatAny(a.a, value)
}

// todo: take equator

func ThatAny(a assert_type.AssertType, value any) *asserter.Any {
	return asserter.NewAny(a, equator.NewReflectDeepEquator[any](), value)
}
