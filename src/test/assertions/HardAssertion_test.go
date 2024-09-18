package assertions

import (
	. "github.com/sku0x20/assertG/src/pkg/assertions"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	. "github.com/sku0x20/assertG/src/test/api"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_HardAssertion(tm *testing.T) {
	r := runner.NewTestsRunner[*instances](tm, initE)
	r.Add(printsMsg)
	r.Add(existsOnFailure)
	r.Run()
}

func initE(t *testing.T) *instances {
	fakeT := NewFakeT(t)
	ha := NewHardAssertion(fakeT)
	return &instances{
		fakeT,
		ha,
	}
}

func printsMsg(t *testing.T, i *instances) {
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	i.ha.FailWith(msg)
	if i.ft.Error != msg.ToString() {
		t.Fatalf("expected '%s', got '%s'", msg.ToString(), i.ft.Error)
	}
}

func existsOnFailure(t *testing.T, i *instances) {
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	i.ha.FailWith(msg)
	i.ft.AssertIsFatal()
}

type instances struct {
	ft *FakeT
	ha *HardAssertion
}
