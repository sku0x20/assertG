package assertions

import (
	. "assertG/src/pkg/assertions"
	"assertG/src/pkg/message"
	"assertG/src/pkg/message/verbs"
	. "assertG/src/test/types"
	"testing"
)

//func Test_PrintMsg(t *testing.T) {
//	//t.Fatalf("todo")
//}

func Test_ExistsOnFailure(t *testing.T) {
	ft := NewFakeT(t)
	ha := NewHardAssertion(ft)
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	ha.FailWith(msg)
	ft.AssertIsFatal()
}
