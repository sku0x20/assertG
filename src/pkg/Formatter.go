package pkg

import "fmt"

type Formatter struct {
}

func (f *Formatter) FormatMismatch(actual any, expected any) string {
	return fmt.Sprintf("Expected:\n%v\nActual:\n%v", expected, actual)
}

func (f *Formatter) FormatMatch(actual any) string {
	return fmt.Sprintf("Not Expected:\n%v", actual)
}

func (f *Formatter) Format(actual any, assertion string, expected any) string {
	return fmt.Sprintf("Expected:\n%v\n%s:\n%v", actual, assertion, expected)
}

func NewFormatter() *Formatter {
	return &Formatter{}
}
