package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestPanicsIfTIsNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("GetT should have panicked")
		}
	}()
	pkg.GetT()
}

func TestClean(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("GetT should have panicked")
		}
	}()
	pkg.RegisterT(t)
	pkg.CleanT()
	pkg.GetT()
	t.Fatalf("GetT should have panicked")
}

func TestEnableAsserts(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("GetT should have panicked")
		}
	}()
	defer func() {
		if pkg.GetT() != nil {
			t.Fatalf("cleanup failed")
		}
	}()
	defer pkg.EnableAsserts(t)()

	if pkg.GetT() != t {
		t.Fatalf("unable to get T")
	}
}
