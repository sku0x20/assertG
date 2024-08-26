package assert

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/asserter"
)

func (c *CaptureT) ThatAny(actual any) *asserter.AnyAsserter {
	return c.AssertThatAny(actual)
}

//goland:noinspection GoNameStartsWithPackageName
func (c *CaptureT) AssertThatAny(actual any) *asserter.AnyAsserter {
	h := pkg.NewAssertHelper(c.GetT(), actual)
	return asserter.NewAnyAsserter(h)
}
