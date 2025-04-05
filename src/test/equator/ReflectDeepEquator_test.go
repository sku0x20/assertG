package equator

import (
	"github.com/sku0x20/assertG/src/main/equator"
	"testing"
)

func Test_ReflectDeepEquator(t *testing.T) {
	type sa struct {
		a int
		b string
		c []string
	}
	a := sa{10, "10", []string{"10"}}
	a2 := sa{10, "10", []string{"10"}}

	b := struct {
		a int
		b string
		c []string
	}{10, "10", []string{"20"}}

	c := struct {
		a int
		b int
		c []string
	}{10, 10, []string{"10"}}

	rde := equator.NewReflectDeepEquator[sa]()
	if !rde.AreEqual(a, a2) {
		t.Fatalf("should be equal")
	}
	if rde.AreEqual(a, b) {
		t.Fatalf("should not be equal")
	}
	if rde.AreEqual(a, c) {
		t.Fatalf("should not be equal")
	}
	if rde.AreEqual(b, c) {
		t.Fatalf("should not be equal")
	}
}
