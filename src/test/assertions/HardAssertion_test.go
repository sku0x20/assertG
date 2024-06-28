package assertions

import (
	. "assertG/src/pkg/assertions"
	"assertG/src/pkg/message"
	"assertG/src/pkg/message/verbs"
	. "assertG/src/test/types"
	"testing"
)

func Test_PrintMsg(t *testing.T) {
	ft, ha := createInstances(t)
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	ha.FailWith(msg)
	if ft.Error != msg.ToString() {
		t.Fatalf("expected '%s', got '%s'", msg.ToString(), ft.Error)
	}
}

func Test_ExistsOnFailure(t *testing.T) {
	ft, ha := createInstances(t)
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	ha.FailWith(msg)
	ft.AssertIsFatal()
}

func createInstances(t *testing.T) (*FakeT, *HardAssertion) {
	ft := NewFakeT(t)
	ha := NewHardAssertion(ft)
	return ft, ha
}
