package assert

import (
	"assertG/src/pkg"
	"assertG/src/pkg/asserter"
)

func AssertThatAny(actual any) *asserter.AnyAsserter {
	return createAnyAsserter(actual)
}

func ThatAny(actual any) *asserter.AnyAsserter {
	return createAnyAsserter(actual)
}

func createAnyAsserter(actual any) *asserter.AnyAsserter {
	h := pkg.NewAssertHelper(pkg.GetT(), actual)
	return asserter.NewAnyAsserter(h)
}
