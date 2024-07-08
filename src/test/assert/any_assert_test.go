package assert

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"testing"
)

func TestAssertThatAny(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var a any = assert.AssertThatAny(10)
	_, ok := a.(*asserter.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThatAny(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var a any = assert.ThatAny(10)
	_, ok := a.(*asserter.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
