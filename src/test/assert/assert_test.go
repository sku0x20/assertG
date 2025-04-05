package assert

import (
	assertP "github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"testing"
)

func Test_Assert(t *testing.T) {
	var assert any
	assert = assertP.NewAssert(assert_type.NewSoftAssert(t))
	casted, ok := assert.(*assertP.Assert)
	if !ok {
		t.Fatalf("err")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
