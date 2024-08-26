package assert

import (
	"github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/asserter"
	"testing"
)

func TestAssertThatAny(t *testing.T) {
	c := assert.NewCaptureT(t)
	var a any = c.AssertThatAny(10)
	casted, ok := a.(*asserter.AnyAsserter)
	if casted == nil {
		t.Fatalf("casted is nil")
	}
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThatAny(t *testing.T) {
	c := assert.NewCaptureT(t)
	var a any = c.ThatAny(10)
	casted, ok := a.(*asserter.AnyAsserter)
	if casted == nil {
		t.Fatalf("casted is nil")
	}
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
