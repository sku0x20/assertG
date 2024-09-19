package assertion

import (
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"github.com/sku0x20/assertG/src/pkg/message"
	"github.com/sku0x20/assertG/src/pkg/message/verbs"
	"testing"
)

func Test_Soft(t *testing.T) {
	s := assertion.NewSoft(t)
	if s.Failed() {
		t.Fatalf("should not be in failed state")
	}
	s.FailWith(message.Expected().
		Value(10).
		Verb(verbs.ToEqual).
		Value(20),
	)
	if !s.Failed() {
		t.Fatalf("should have failed")
	}
}
