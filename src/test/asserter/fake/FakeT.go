package fake

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

func (ft *FakeT) AssertIsFatal(t *testing.T, message string) {
	if !ft.isFatal {
		t.Fatalf("should be fatal")
	}
	if ft.error != message {
		t.Fatalf("invalid message = '%v'", ft.error)
	}
}

func (ft *FakeT) AssertNotFatal(t *testing.T) {
	if ft.isFatal {
		t.Fatalf("should not be fatal")
	}
}

func (ft *FakeT) Reset() {
	ft.isFatal = false
	ft.error = ""
}
