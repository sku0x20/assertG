package assert

import (
	"assertG/src/pkg"
	"testing"
)

func TestCorrectAsserter(t *testing.T) {
	pkg.RegisterT(t)
	// todo:
}

// support assert.That(12).isEqualTo(12) also
