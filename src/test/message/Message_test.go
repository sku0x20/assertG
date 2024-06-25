package message

import (
	"assertG/src/pkg/message"
	"testing"
)

func Test_Message(t *testing.T) {
	str := message.NewMessage().
		NewLine().
		Verb("Expected").
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

func Test_OnlyVerbs(t *testing.T) {
	str := message.NewMessage().
		NewLine().
		Verb("Expected").
		Verb("ToBeNil").
		Verb("ButWas").
		ToString()
	expected := `
Expected
ToBeNil
ButWas
`
	if str != expected {
		t.Fatalf("failure; msg = %s", str)
	}
}

func Test_OnlyValues(t *testing.T) {
	str := message.NewMessage().
		NewLine().
		Value("a").
		Value("b").
		ToString()
	expected := `
a
b
`
	if str != expected {
		t.Fatalf("failure; msg = %s", str)
	}
}
