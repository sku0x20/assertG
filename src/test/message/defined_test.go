package message

import (
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/assertG/src/main/message/verbs"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Defined(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(expected)
	r.Run()
}

func expected(t *testing.T, _ any) {
	e := message.Expected()
	msg := message.NewMessage().
		NewLine().
		Verb(verbs.Expected)
	if e.ToString() != msg.ToString() {
		t.Fatalf("Expected: %s, got: %s", e.ToString(), msg.ToString())
	}
}
