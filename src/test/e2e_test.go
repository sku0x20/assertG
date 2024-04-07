package test

import (
	"assertG/src/pkg"
	"assertG/src/pkg/assert"
	"testing"
)

func TestUsage(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	assert.That(10).IsEqualTo(10)
}
