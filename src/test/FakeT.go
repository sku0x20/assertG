package test

import "fmt"

type FakeT struct {
	isFatal bool
	error   string
}

func (t *FakeT) Fatalf(format string, args ...any) {
	t.isFatal = true
	t.error = fmt.Sprintf(format, args...)
}
