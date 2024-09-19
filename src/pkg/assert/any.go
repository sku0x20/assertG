package assert

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/pkg/assertion"
)

func (a *Assert) ThatAny(value any) {

}

func ThatAny(a assertion.Assertion) *asserter.Any {
	return asserter.NewAny(a)
}
