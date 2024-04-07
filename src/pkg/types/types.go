package types

type T interface {
	Fatalf(format string, args ...any)
}

type Equals interface {
	Equals(other any) bool
}
