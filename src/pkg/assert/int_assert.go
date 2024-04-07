package assert

import "assertG/src/pkg"

func AssertThat(actual int) *pkg.Asserter {
	return createAsserter(actual)
}

func That(actual int) *pkg.Asserter {
	return createAsserter(actual)
}

func createAsserter(actual int) *pkg.Asserter {
	return pkg.NewAsserter(pkg.GetT(), actual)
}
