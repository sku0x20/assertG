package assert

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
)

func ThatAny(actual any) *asserter.AnyAsserter {
	return AssertThatAny(actual)
}

//goland:noinspection GoNameStartsWithPackageName
func AssertThatAny(actual any) *asserter.AnyAsserter {
	//h := pkg.NewAssertHelper(pkg.GetT(), actual)
	//return asserter.NewAnyAsserter(h)
	return nil
}
