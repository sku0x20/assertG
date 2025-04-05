package assert

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/assertG/src/main/registry"
)

func ThatSlice[T any](value []T) *asserter.Slice[T] {
	at := registry.GlobalRegistryGetAssertType()
	return ThatSliceWith(at, equator.NewReflectDeepEquator[T](), value)
}

func ThatSliceWith[T any](
	a assert_type.AssertType,
	elemEquator equator.Equator[T],
	value []T,
) *asserter.Slice[T] {
	return asserter.NewSlice(a, elemEquator, value)
}
