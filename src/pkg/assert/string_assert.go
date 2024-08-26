package assert

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/asserter"
)

func (c *CaptureT) ThatString(actual string) *asserter.StringAsserter {
	return c.AssertThatString(actual)
}

//goland:noinspection GoNameStartsWithPackageName
func (c *CaptureT) AssertThatString(actual string) *asserter.StringAsserter {
	h := pkg.NewAssertHelper(c.GetT(), actual)
	return asserter.NewStringAsserter(h)
}
