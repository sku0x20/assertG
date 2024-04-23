package pkg

import "fmt"

type Formatter struct {
}

func (f *Formatter) FormatError(actual string, expected string) string {
	return fmt.Sprintf("Expected:\n%s\nActual:\n%s", expected, actual)
}

func NewFormatter() *Formatter {
	return &Formatter{}
}
