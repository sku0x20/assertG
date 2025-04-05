package message

import (
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Message(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(fullMessage)
	r.Add(onlyVerbs)
	r.Add(onlyValues)
	r.Run()
}

func fullMessage(t *testing.T, _ any) {
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
	e := `
Expected
a
ToEqual
20
ToHaveLength
21
ButWas
12
`
	if str != e {
		t.Fatalf("failure; msg = %s", str)
	}
}

func onlyVerbs(t *testing.T, _ any) {
	str := message.NewMessage().
		NewLine().
		Verb("Expected").
		Verb("ToBeNil").
		Verb("ButWas").
		ToString()
	e := `
Expected
ToBeNil
ButWas
`
	if str != e {
		t.Fatalf("failure; msg = %s", str)
	}
}

func onlyValues(t *testing.T, _ any) {
	str := message.NewMessage().
		NewLine().
		Value("a").
		Value("b").
		ToString()
	e := `
a
b
`
	if str != e {
		t.Fatalf("failure; msg = %s", str)
	}
}
