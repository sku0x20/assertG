package assert

import (
	"assertG/src/pkg"
	"assertG/src/pkg/assert"
	"testing"
)

func TestAssertThat_String(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.AssertThatString("string assert")
	_, ok := asserter.(*pkg.StringAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}

func TestThat_String(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	var asserter any = assert.ThatString("string assert")
	_, ok := asserter.(*pkg.StringAsserter)
	if !ok {
		t.Fatalf("invalid asserter type")
	}
}
