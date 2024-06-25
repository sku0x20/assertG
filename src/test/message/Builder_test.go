package message

import (
	"assertG/src/pkg/message"
	"testing"
)

func Test_MultipleMessages(t *testing.T) {
	str := message.Expected().
		Value("a").
		Verb("ToEqual").
		Value("20").
		Verb("ToHaveLength").
		Value("21").
		Verb("ButWas").
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

func Test_ContinuousMessages(t *testing.T) {
	str := message.Expected().
		Verb("ToBeNil").
		Verb("ButWas").
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

func Test_EndingWithMessage(t *testing.T) {
	str := message.Expected().
		Value("a").
		Verb("ToBeNil").
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
