package assert

import (
	"github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"testing"
)

func TestAssertThat_String(t *testing.T) {
	var a any = assert.AssertThatString("string assert")
	casted, ok := a.(*asserter.StringAsserter)
	if casted == nil {
		t.Fatalf("casted is nil")
	}
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThat_String(t *testing.T) {
	var a any = assert.ThatString("string assert")
	casted, ok := a.(*asserter.StringAsserter)
	if casted == nil {
		t.Fatalf("casted is nil")
	}
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
