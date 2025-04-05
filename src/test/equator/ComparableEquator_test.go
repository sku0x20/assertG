package equator

import (
	"github.com/sku0x20/assertG/src/main/equator"
	"testing"
)

func Test_ComparableEquator(t *testing.T) {
	a := 10
	a2 := 10

	b := struct {
		a int
		b string
		c []string
	}{10, "10", []string{"20"}}

	rde := equator.NewComparableEquator[int]()
	if !rde.AreEqual(a, a2) {
		t.Fatalf("should be equal")
	}
	if rde.AreEqual(a, b) {
		t.Fatalf("should not be equal")
	}
}
