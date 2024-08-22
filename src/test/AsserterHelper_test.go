package test

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/test/api"
	"testing"
)

func TestFormatter(t *testing.T) {
	helper := pkg.NewAssertHelper(t, 10)
	formatter := helper.Formatter()
	str := formatter.ToString()
	if str != "\nExpected\n10\n" {
		t.Fatalf("failure; msg = %s", str)
	}
}

func TestError(t *testing.T) {
	ft := api.NewFakeT(t)
	helper := pkg.NewAssertHelper(ft, 10)
	formatter := helper.Formatter().
		Value("10")
	helper.Error(formatter)
	ft.AssertIsFatal()
	if formatter.ToString() != ft.Error {
		t.Fatalf("failure; msg = %s\n%s", ft.Error, formatter.ToString())
	}
}
