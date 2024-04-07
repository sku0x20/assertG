package assert

import "assertG/src/pkg"

func AssertThat(actual int) *pkg.AnyAsserter {
	return createAsserter(actual)
}

func That(actual int) *pkg.AnyAsserter {
	return createAsserter(actual)
}

func createAsserter(actual int) *pkg.AnyAsserter {
	return pkg.NewAnyAsserter(pkg.GetT(), actual)
}
