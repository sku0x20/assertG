package assert

import "assertG/src/pkg"

func AssertThatAny(actual any) *pkg.AnyAsserter {
	return createAnyAsserter(actual)
}

func ThatAny(actual any) *pkg.AnyAsserter {
	return createAnyAsserter(actual)
}

func createAnyAsserter(actual any) *pkg.AnyAsserter {
	return pkg.NewAnyAsserter(pkg.GetT(), actual)
}
