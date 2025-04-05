package test

import (
	assertP "github.com/sku0x20/assertG/src/pkg/assert"
	"github.com/sku0x20/assertG/src/pkg/assert_type"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_WithoutRunner(t *testing.T) {
	assert := assertP.NewAssert(assert_type.NewHardAssert(t))
	assert.ThatAny(10).
		IsNotNil().
		IsEqualTo(10)
}

func initE(t *testing.T) *assertP.Assert {
	return assertP.NewAssert(assert_type.NewHardAssert(t))
}

func Test_WithRunner(t *testing.T) {
	r := runner.NewTestsRunner[*assertP.Assert](t, initE)
	r.Add(makeAssert)
	r.Run()
}

func makeAssert(_ *testing.T, assert *assertP.Assert) {
	assert.ThatAny(10).
		IsNotNil().
		IsEqualTo(10)
}
