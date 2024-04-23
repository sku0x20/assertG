package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestFormatErrorInt(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatError(1, 2)
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatErrorString(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatError("1", "2")
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}
