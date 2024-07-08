package assert

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"testing"
)

func TestAssertThat_String(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var a any = assert.AssertThatString("string assert")
	_, ok := a.(*asserter.StringAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThat_String(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var a any = assert.ThatString("string assert")
	_, ok := a.(*asserter.StringAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
