package pkg

import "fmt"

type Formatter struct {
}

func (f *Formatter) FormatError(actual any, expected any) string {
	return fmt.Sprintf("Expected:\n%v\nActual:\n%v", expected, actual)
}

func NewFormatter() *Formatter {
	return &Formatter{}
}
