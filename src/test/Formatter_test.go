package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestFormatError(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatError("a", "b")
	if str != "Expected:\nb\nActual:\na" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}
