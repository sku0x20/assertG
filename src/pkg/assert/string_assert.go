package assert

import (
	"assertG/src/pkg"
	"assertG/src/pkg/asserter"
)

func ThatString(actual string) *asserter.StringAsserter {
	return AssertThatString(actual)
}

//goland:noinspection GoNameStartsWithPackageName
func AssertThatString(actual string) *asserter.StringAsserter {
	h := pkg.NewAssertHelper(pkg.GetT(), actual)
	return asserter.NewStringAsserter(h)
}
