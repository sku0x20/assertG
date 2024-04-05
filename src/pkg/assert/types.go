package assert

type T interface {
	Fatalf(format string, args ...any)
}
