package equator

type Equator[T any] interface {
	compare(a T, b any) bool
}
