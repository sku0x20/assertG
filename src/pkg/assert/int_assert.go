package assert

import "assertG/src/pkg"

func AssertThat(actual int) *pkg.IntAsserter {
	return createAsserter(actual)
}

func That(actual int) *pkg.IntAsserter {
	return createAsserter(actual)
}

func createAsserter(actual int) *pkg.IntAsserter {
	return pkg.NewAsserter(pkg.GetT(), actual)
}
