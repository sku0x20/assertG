package assert

import (
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"github.com/sku0x20/assertG/src/pkg/asserter"
)

func (a *Assert) ThatString(value string) *asserter.String {
	return ThatString(a.a, value)
}

func ThatString(a assert_type.AssertType, value string) *asserter.String {
	return asserter.NewString(a, value)
}
