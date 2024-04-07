package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestIsNotNil(t *testing.T) {
	ba := pkg.NewBoolAsserter(t, true)
	ba.IsNotNil()
}
