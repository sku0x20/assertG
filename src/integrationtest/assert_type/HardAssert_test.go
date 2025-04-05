package assert_type

import (
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"strings"
	"testing"
)

func Test_Hard(t *testing.T) {
	if runReal() {
		hardReal(t)
	} else {
		output := spawnAndRun(funcName())
		lines := splitAndTrimLines(output)
		if !strings.Contains(lines[3], "Expected") {
			t.Fatalf("should print 'Expected'")
		}
		if !strings.Contains(lines[4], "ToBeNil") {
			t.Fatalf("should print 'ToBeNil'")
		}
		if lines[5] != "FAIL" {
			t.Fatalf("should have failed")
		}
		if len(lines) >= 8 {
			if lines[7] == "shouldn't be printed" {
				t.Fatalf("printing `shouldn't be printed`")
			}
		}
	}
}

func hardReal(t *testing.T) {
	r := runner.NewTestsRunner[*assert_type.HardAssert](t, initE)
	r.Add(printsMsg)
	r.Run()
}

func initE(t *testing.T) *assert_type.HardAssert {
	return assert_type.NewHardAssert(t)
}

func printsMsg(t *testing.T, h *assert_type.HardAssert) {
	msg := message.Expected().
		Verb(verbs.ToBeNil)
	h.FailWith(msg)
	t.Fatalf("shouldn't be printed")
}
