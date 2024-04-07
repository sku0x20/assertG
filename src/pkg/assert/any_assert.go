package assert

import "assertG/src/pkg"

func AssertThat(actual any) *pkg.AnyAsserter {
	return createAsserter(actual)
}

func That(actual any) *pkg.AnyAsserter {
	return createAsserter(actual)
}

func createAsserter(actual any) *pkg.AnyAsserter {
	return pkg.NewAnyAsserter(pkg.GetT(), actual)
}
