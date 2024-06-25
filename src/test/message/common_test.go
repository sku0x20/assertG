package message

import (
	"assertG/src/pkg/message"
	"assertG/src/pkg/message/verbs"
	"testing"
)

func Test_Expected(t *testing.T) {
	expected := message.Expected()
	msg := message.NewMessage().
		NewLine().
		Verb(verbs.Expected)
	if expected.ToString() != msg.ToString() {
		t.Errorf("Expected: %s, got: %s", expected.ToString(), msg.ToString())
	}
}
