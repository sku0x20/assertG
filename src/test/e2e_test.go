package test

import (
	"github.com/sku0x20/assertG/src/pkg"
	"github.com/sku0x20/assertG/src/pkg/assert"
	"testing"
)

func TestUsage(t *testing.T) {
	defer pkg.EnableAsserts(t)()
	assert.ThatAny(10).IsNotNil().IsEqualTo(10)
}
