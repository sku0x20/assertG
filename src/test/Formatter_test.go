package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestFormatMismatchInt(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatMismatch(1, 2)
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatMismatchString(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatMismatch("1", "2")
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatMismatch(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatMismatch("1", 2)
	if str != "Expected:\n2\nActual:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatMatchString(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatMatch("1")
	if str != "Not Expected:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}

func TestFormatMatchInt(t *testing.T) {
	formatter := pkg.NewFormatter()
	str := formatter.FormatMatch(1)
	if str != "Not Expected:\n1" {
		t.Errorf("wrong formatted string; \n'%s'\n", str)
	}
}
