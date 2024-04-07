package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestBoolIsNotNil(t *testing.T) {
	ba := pkg.NewBoolAsserter(t, true)
	ba.IsNotNil()
}
