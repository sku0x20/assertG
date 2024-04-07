package assert

import "assertG/src/pkg"

func AssertThat(actual int) *pkg.Asserter {
	t := pkg.GetT()
	return pkg.NewAsserter(t, actual)
}
