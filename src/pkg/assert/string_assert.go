package assert

import "assertG/src/pkg"

func AssertThatString(actual string) *pkg.StringAsserter {
	return createStringAsserter(actual)
}

func ThatString(actual string) *pkg.StringAsserter {
	return createStringAsserter(actual)
}

func createStringAsserter(actual string) *pkg.StringAsserter {
	return pkg.NewStringAsserter(pkg.GetT(), actual)
}
