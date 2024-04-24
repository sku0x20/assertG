package asserter

import (
	"assertG/src/pkg"
	"assertG/src/pkg/asserter"
	"assertG/src/test/asserter/fake"
	"testing"
)

func TestAnyEquals(t *testing.T) {
	ft := &fake.FakeT{}
	ga := asserter.NewAnyAsserter(ft, pkg.NewFormatter(), 10)
	ga.IsEqualTo(30)
	ft.AssertIsFatal(t)
	ft.Reset()
	ga.IsEqualTo(10)
	ft.AssertNotFatal(t)
}

func TestAnyEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := &fake.FakeT{}
	equalFalse := &fake.FakeEquals{false}
	equalTrue := &fake.FakeEquals{true}
	ga1 := asserter.NewAnyAsserter(ft, pkg.NewFormatter(), equalFalse)
	ga1.IsEqualTo(equalTrue)
	ft.AssertIsFatal(t)
	ft.Reset()
	ga2 := asserter.NewAnyAsserter(ft, pkg.NewFormatter(), equalTrue)
	ga2.IsEqualTo(equalFalse)
	ft.AssertNotFatal(t)
}

func TestAnyNotEquals(t *testing.T) {
	ft := &fake.FakeT{}
	ga := asserter.NewAnyAsserter(ft, pkg.NewFormatter(), 10)
	ga.IsNotEqualTo(10)
	ft.AssertIsFatal(t)
	ft.Reset()
	ga.IsNotEqualTo(20)
	ft.AssertNotFatal(t)
}

func TestAnyNotEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := &fake.FakeT{}
	fE1 := &fake.FakeEquals{false}
	fE2 := &fake.FakeEquals{true}
	ga1 := asserter.NewAnyAsserter(ft, pkg.NewFormatter(), fE2)
	ga1.IsNotEqualTo(fE1)
	ft.AssertIsFatal(t)
	ft.Reset()
	ga2 := asserter.NewAnyAsserter(ft, pkg.NewFormatter(), fE1)
	ga2.IsNotEqualTo(fE2)
	ft.AssertNotFatal(t)
}

func TestAnyNil(t *testing.T) {
	ft := &fake.FakeT{}
	actual := 10
	asserter.NewAnyAsserter(ft, pkg.NewFormatter(), &actual).IsNil()
	ft.AssertIsFatal(t)
	ft.Reset()
	asserter.NewAnyAsserter(ft, pkg.NewFormatter(), nil).IsNil()
	ft.AssertNotFatal(t)
}

func TestAnyNotNil(t *testing.T) {
	ft := &fake.FakeT{}
	asserter.NewAnyAsserter(ft, pkg.NewFormatter(), nil).IsNotNil()
	ft.AssertIsFatal(t)
	ft.Reset()
	actual := 10
	asserter.NewAnyAsserter(ft, pkg.NewFormatter(), &actual).IsNotNil()
	ft.AssertNotFatal(t)
}
