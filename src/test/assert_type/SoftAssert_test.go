package assert_type

import (
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/message"
	"github.com/sku0x20/assertG/src/main/message/verbs"
	"testing"
)

func Test_Soft(t *testing.T) {
	s := assert_type.NewSoftAssert(t)
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
