package assert

import (
	assertP "github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/assertion"
	"testing"
)

func Test_Assert(t *testing.T) {
	var assert any
	assert = assertP.NewAssert(assertion.NewSoft(t))
	casted, ok := assert.(*assertP.Assert)
	if !ok {
		t.Fatalf("err")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
