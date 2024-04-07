package test

import (
	"fmt"
	"testing"
)

type FakeT struct {
	isFatal bool
	error   string
}

func (ft *FakeT) Fatalf(format string, args ...any) {
	ft.isFatal = true
	ft.error = fmt.Sprintf(format, args...)
}

func (ft *FakeT) assertIsFatal(t *testing.T, message string) {
	if !ft.isFatal {
		t.Fatalf("should be fatal")
	}
	if ft.error != message {
		t.Fatalf("invalid message = '%v'", ft.error)
	}
}
