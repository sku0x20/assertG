package assertions

import (
	. "assertG/src/pkg/assertions"
	"assertG/src/pkg/message"
	"assertG/src/pkg/message/verbs"
	. "assertG/src/test/types"
	"testing"
)

func Test_HardAssertion(t *testing.T) {
	ft := NewFakeT(t)
	ha := NewHardAssertion(ft)

	t.Run("Print_Msg", func(t *testing.T) {
		msg := message.Expected().
			Verb(verbs.ToBeNil)
		ha.FailWith(msg)
		if ft.Error != msg.ToString() {
			t.Fatalf("expected '%s', got '%s'", msg.ToString(), ft.Error)
		}
	})

	t.Run("ExistsOnFailure", func(t *testing.T) {
		msg := message.Expected().
			Verb(verbs.ToBeNil)
		ha.FailWith(msg)
		ft.AssertIsFatal()
	})
}
