package assert

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/assertG/src/main/registry"
)

// methods cannot have type parameters
// using global registry for AssertType

func That[T any](value T) *asserter.Generic[T] {
	at := registry.GlobalRegistryGetAssertType()
	return ThatWith(at, equator.NewReflectDeepEquator[T](), value)
}

func ThatWith[T any](
	at assert_type.AssertType,
	equator equator.Equator[T],
	value T,
) *asserter.Generic[T] {
	return asserter.NewGeneric(at, equator, value)
}
