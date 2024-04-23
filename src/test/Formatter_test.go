package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestFormatDifferentInt(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatError(1, 2)
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatDifferentString(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatError("1", "2")
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatDifferent(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatError("1", 2)
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatSame(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatErrorSame("1")
	if str != "Not Expected:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}
