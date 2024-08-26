package assert

import (
	"github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"testing"
)

func TestAssertThat_String(t *testing.T) {
	c := assert.NewCaptureT(t)
	var a any = c.AssertThatString("string assert")
	casted, ok := a.(*asserter.StringAsserter)
	if casted == nil {
		t.Fatalf("casted is nil")
	}
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThat_String(t *testing.T) {
	c := assert.NewCaptureT(t)
	var a any = c.ThatString("string assert")
	casted, ok := a.(*asserter.StringAsserter)
	if casted == nil {
		t.Fatalf("casted is nil")
	}
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
