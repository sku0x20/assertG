package format

import (
	"assertG/src/pkg/format"
	"testing"
)

func TestBuildsString1(t *testing.T) {
	str := format.Expected().
		Value("a").
		Message("ToEqual").
		Value("20").
		Message("ToHaveLength").
		Value("21").
		Message("ButWas").
		Value("12").
		ToString()
	expected := `
Expected
a
ToEqual
20
ToHaveLength
21
ButWas
12
`
	if str != expected {
		t.Fatalf("failure; msg = %s", str)
	}
}

func TestBuildString2(t *testing.T) {
	str := format.Expected().
		Message("ToBeNil").
		Message("ButWas").
		Value("a").
		ToString()
	expected := `
Expected
ToBeNil
ButWas
a
`
	if str != expected {
		t.Fatalf("failure; msg = %s", str)
	}
}

func TestBuildString3(t *testing.T) {
	str := format.Expected().
		Value("a").
		Message("ToBeNil").
		ToString()
	expected := `
Expected
a
ToBeNil
`
	if str != expected {
		t.Fatalf("failure; msg = %s", str)
	}
}
