package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestTHolder(t *testing.T) {
	pkg.RegisterT(t)
	if pkg.GetT() != t {
		t.Fatalf("unable to get T")
	}
	pkg.CleanT()
	if pkg.GetT() != nil {
		t.Fatalf("cleanup failed")
	}
}
