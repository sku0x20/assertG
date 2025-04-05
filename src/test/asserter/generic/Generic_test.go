package generic

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/assertG/src/main/equator"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func init_G(t *testing.T) *assert_type.SoftAssert {
	return assert_type.NewSoftAssert(t)
}

func Test_Generic(t *testing.T) {
	r := runner.NewTestsRunner[*assert_type.SoftAssert](t, init_G)
	r.Add(equalFail)
	r.Add(equalPass)
	r.Add(customEquator)
	r.Add(nilPass)
	r.Add(nilFail)
	r.Add(notNilPass)
	r.Add(notNilFail)
	r.Add(nilSlice)
	r.Run()
}

func equalFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewGeneric(s, equator.NewReflectDeepEquator[int](), 30)
	a = a.IsEqualTo("something")
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func equalPass(t *testing.T, s *assert_type.SoftAssert) {
	type st struct {
		s *string
	}
	s1 := "something"
	s2 := "something"
	a := asserter.NewGeneric(s, equator.NewReflectDeepEquator[*st](), &st{&s1})
	a = a.IsEqualTo(&st{&s2})
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

type FakeEquator struct {
}

//goland:noinspection GoUnusedParameter
func (f *FakeEquator) AreEqual(a int, b any) bool {
	return true
}

func customEquator(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewGeneric(s, &FakeEquator{}, 10)
	a = a.IsEqualTo(20)
	if s.Failed() {
		t.Fatalf("should have passed")
	}
}

func nilPass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewGeneric(s, equator.NewReflectDeepEquator[any](), nil)
	a = a.IsNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func nilFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewGeneric(s, equator.NewReflectDeepEquator[int](), 10)
	a = a.IsNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func notNilPass(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewGeneric(s, equator.NewReflectDeepEquator[int](), 10)
	a = a.IsNotNil()
	if s.Failed() {
		t.Fatalf("should not have failed")
	}
}

func notNilFail(t *testing.T, s *assert_type.SoftAssert) {
	a := asserter.NewGeneric(s, equator.NewReflectDeepEquator[any](), nil)
	a = a.IsNotNil()
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}

func rns() []string {
	var ns []string
	return ns
}

func nilSlice(t *testing.T, sa *assert_type.SoftAssert) {
	a := asserter.NewGeneric(sa, equator.NewReflectDeepEquator[[]string](), rns())
	a.IsNil()
	if !sa.Failed() {
		t.Fatalf("should have failed")
	}
}
