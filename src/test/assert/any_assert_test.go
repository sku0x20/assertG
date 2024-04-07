package assert

import (
	"assertG/src/pkg"
	"assertG/src/pkg/assert"
	"testing"
)

func TestAssertThatAny(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.AssertThat(10)
	_, ok := asserter.(*pkg.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThat_Any(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.That(10)
	_, ok := asserter.(*pkg.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
