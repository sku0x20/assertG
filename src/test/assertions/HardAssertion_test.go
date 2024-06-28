package assertions

import (
	. "assertG/src/pkg/assertions"
	"assertG/src/pkg/message"
	"assertG/src/pkg/message/verbs"
	. "assertG/src/test/types"
	"testing"
)

func Test_HardAssertion(t *testing.T) {
	for tName, tCase := range tests {
		t.Run(tName, func(t *testing.T) {
			ft := NewFakeT(t)
			ha := NewHardAssertion(ft)
			tCase(t, ft, ha)
		})
	}
}

var tests = map[string]testCase{
	"PrintMsg":        PrintMsg,
	"ExistsOnFailure": ExistsOnFailure,
}

type testCase func(t *testing.T, ft *FakeT, ha *HardAssertion)

func PrintMsg(
	t *testing.T,
	ft *FakeT,
	ha *HardAssertion,
) {
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	ha.FailWith(msg)
	if ft.Error != msg.ToString() {
		t.Fatalf("expected '%s', got '%s'", msg.ToString(), ft.Error)
	}
}

func ExistsOnFailure(
	t *testing.T,
	ft *FakeT,
	ha *HardAssertion,
) {
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	ha.FailWith(msg)
	ft.AssertIsFatal()
}
