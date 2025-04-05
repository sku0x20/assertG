package equator

type Equator[T any] interface {
	AreEqual(a T, b any) bool
}
