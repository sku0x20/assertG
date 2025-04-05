package assert

import (
	assertP "github.com/sku0x20/assertG/src/main/assert"
	"github.com/sku0x20/assertG/src/main/assert_type"
	"github.com/sku0x20/assertG/src/main/asserter"
	"github.com/sku0x20/gRunner/src/pkg/runner"
	"testing"
)

func Test_Generic(t *testing.T) {
	r := runner.NewTestsRunnerEmptyInit[any](t)
	r.Add(viaPackage_G)
	r.Add(viaVariable_G)
	r.Run()
}

func viaPackage_G(t *testing.T, e any) {
	var a any = assertP.That(assert_type.NewSoftAssert(t), "some val")
	casted, ok := a.(*asserter.Generic[string])
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}

func viaVariable_G(t *testing.T, e any) {
	assert := assertP.NewAssert(assert_type.NewSoftAssert(t))
	var a any = assert.That("some val")
	casted, ok := a.(*asserter.Generic[any])
	if !ok {
		t.Fatalf("unable to cast")
	}
	if casted == nil {
		t.Fatalf("casted is nil")
	}
}
