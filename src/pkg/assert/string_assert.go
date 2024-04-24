package assert

import (
	"assertG/src/pkg"
	"assertG/src/pkg/asserter"
)

func AssertThatString(actual string) *asserter.StringAsserter {
	return createStringAsserter(actual)
}

func ThatString(actual string) *asserter.StringAsserter {
	return createStringAsserter(actual)
}

func createStringAsserter(actual string) *asserter.StringAsserter {
	return asserter.NewStringAsserter(pkg.GetT(), pkg.NewFormatter(), actual)
}
