package types

type T interface {
	Fatalf(format string, args ...any)
}

type Equals interface {
	Equals(t T) bool
}
