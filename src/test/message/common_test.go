package message

import (
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
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
