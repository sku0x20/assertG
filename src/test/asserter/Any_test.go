package asserter

import (
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"github.com/sku0x20/assertG/src/test/assertion"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Any(t *testing.T) {
	r := runner.NewTestsRunner[*assertion.Fake](t, initE)
	r.Add(equalsFail)
	r.Run()
}

func initE(t *testing.T) *assertion.Fake {
	return assertion.NewFake()
}

func equalsFail(t *testing.T, f *assertion.Fake) {
	a := asserter.NewAny(f, 30)
	a.IsEqualTo("something")
	if !f.Failed() {
		t.Fatalf("should have failed")
	}
}

/*


func TestAnyEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := api.NewFakeT(t)
	equalFalse := &api.FakeEquals{false}
	equalTrue := &api.FakeEquals{true}
	h := pkg.NewAssertHelper(ft, equalFalse)
	ga1 := asserter.NewAnyAsserter(h)
	ga1.IsEqualTo(equalTrue)
	ft.AssertIsFatal()
	ft.Reset()
	h2 := pkg.NewAssertHelper(ft, equalTrue)
	ga2 := asserter.NewAnyAsserter(h2)
	ga2.IsEqualTo(equalFalse)
	ft.AssertNotFatal()
}

func TestAnyNotEquals(t *testing.T) {
	ft := api.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, 10)
	ga := asserter.NewAnyAsserter(h)
	ga.IsNotEqualTo(10)
	ft.AssertIsFatal()
	ft.Reset()
	ga.IsNotEqualTo(20)
	ft.AssertNotFatal()
}

func TestAnyNotEqualsCallsEqualsIfPossible(t *testing.T) {
	ft := api.NewFakeT(t)
	fE1 := &api.FakeEquals{false}
	fE2 := &api.FakeEquals{true}
	h := pkg.NewAssertHelper(ft, fE2)
	ga1 := asserter.NewAnyAsserter(h)
	ga1.IsNotEqualTo(fE1)
	ft.AssertIsFatal()
	ft.Reset()
	h2 := pkg.NewAssertHelper(ft, fE1)
	ga2 := asserter.NewAnyAsserter(h2)
	ga2.IsNotEqualTo(fE2)
	ft.AssertNotFatal()
}

func TestAnyNil(t *testing.T) {
	ft := api.NewFakeT(t)
	actual := 10
	h := pkg.NewAssertHelper(ft, &actual)
	asserter.NewAnyAsserter(h).IsNil()
	ft.AssertIsFatal()
	ft.Reset()
	h2 := pkg.NewAssertHelper(ft, nil)
	asserter.NewAnyAsserter(h2).IsNil()
	ft.AssertNotFatal()
}

func TestAnyNotNil(t *testing.T) {
	ft := api.NewFakeT(t)
	h := pkg.NewAssertHelper(ft, nil)
	asserter.NewAnyAsserter(h).IsNotNil()
	ft.AssertIsFatal()
	ft.Reset()
	actual := 10
	h2 := pkg.NewAssertHelper(ft, &actual)
	asserter.NewAnyAsserter(h2).IsNotNil()
	ft.AssertNotFatal()
}
*/
