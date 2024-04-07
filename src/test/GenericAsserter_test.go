package test

import (
	"assertG/src/pkg"
	"testing"
)

func TestGenericIsNotNil(t *testing.T) {
	ga := pkg.NewGenericAsserter(t, 10)
	ga.IsEqualTo(30)
}
