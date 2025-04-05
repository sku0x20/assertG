package assert

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
)

// methods cannot have type parameters
// todo: use global registry for AssertType

func (a *Assert) That(value any) *asserter.Generic[any] {
	return That(a.a, value)
}

// todo: create fn that take equator

func That[T any](a assert_type.AssertType, value T) *asserter.Generic[T] {
	return asserter.NewGeneric(a, equator.NewReflectDeepEquator[T](), value)
}
