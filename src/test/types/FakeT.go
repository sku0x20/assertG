package types

import (
	"fmt"
	"log"
	"testing"
)

type FakeT struct {
	isFatal bool
	Error   string
	t       *testing.T
}

func (ft *FakeT) Fatalf(format string, args ...any) {
	ft.isFatal = true
	ft.Error = fmt.Sprintf(format, args...)
}

func (ft *FakeT) AssertIsFatal() {
	if !ft.isFatal {
		ft.t.Fatalf("should be fatal")
	}
	// todo: remove this
	log.Printf("\n----= failure message =----%s", ft.Error)
}

func (ft *FakeT) AssertNotFatal() {
	if ft.isFatal {
		ft.t.Fatalf("should not be fatal")
	}
}

func (ft *FakeT) Reset() {
	ft.isFatal = false
	ft.Error = ""
}

func NewFakeT(t *testing.T) *FakeT {
	return &FakeT{t: t}
}
