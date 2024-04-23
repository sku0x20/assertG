package fake

import (
	"fmt"
	"strings"
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

func (ft *FakeT) AssertIsFatal(t *testing.T, strs ...string) {
	if !ft.isFatal {
		t.Fatalf("should be fatal")
	}
	for _, s := range strs {
		if !strings.Contains(ft.error, s) {
			t.Fatalf("invalid message = '%v'", ft.error)
		}
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
