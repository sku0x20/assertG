package assert

import (
	"assertG/src/pkg"
	"assertG/src/pkg/assert"
	"testing"
)

func TestAssertThatAny(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.AssertThatAny(10)
	_, ok := asserter.(*pkg.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThatAny(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.ThatAny(10)
	_, ok := asserter.(*pkg.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
