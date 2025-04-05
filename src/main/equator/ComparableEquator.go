package equator

type ComparableEquator[T comparable] struct {
}

func NewComparableEquator[T comparable]() *ComparableEquator[T] {
	return &ComparableEquator[T]{}
}

func (c *ComparableEquator[T]) AreEqual(a T, b any) bool {
	return a == b
}
