package assert

import (
	"assertG/src/pkg"
	"assertG/src/pkg/assert"
	"testing"
)

func TestCorrectAsserter_1(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.AssertThat(10)
	_, ok := asserter.(*pkg.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestCorrectAsserter_2(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.That(10)
	_, ok := asserter.(*pkg.AnyAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
