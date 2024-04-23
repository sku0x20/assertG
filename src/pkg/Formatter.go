package pkg

import "fmt"

type Formatter struct {
}

func (f *Formatter) FormatError(actual any, expected any) string {
	return fmt.Sprintf("Expected:\n%v\nActual:\n%v", expected, actual)
}

func (f *Formatter) FormatErrorSame(actual string) string {
	return fmt.Sprintf("Not Expected:\n%v", actual)
}

func NewFormatter() *Formatter {
	return &Formatter{}
}
