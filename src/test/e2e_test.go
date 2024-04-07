package test

import (
	"assertG/src/pkg"
	"assertG/src/pkg/assert"
	"testing"
)

func TestUsage(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	assert.ThatAny(10).IsNotNil().IsEqualTo(10)
}
